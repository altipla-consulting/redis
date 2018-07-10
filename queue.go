package redis

// import (
// 	"fmt"
// 	"time"

// 	"github.com/altipla-consulting/sentry"
// 	"github.com/golang/protobuf/proto"
// 	"github.com/juju/errors"
// 	"github.com/segmentio/ksuid"
// 	log "github.com/sirupsen/logrus"
// 	"golang.org/x/net/context"
// )

// type Queue struct {
// 	db            *Database
// 	queueKey      string
// 	processingKey string
// 	tasksKey      string
// }

// func (queue *Queue) Push(values ...proto.Message) error {
// 	tasks := []interface{}{}
// 	fields := map[string]interface{}{}
// 	for _, value := range values {
// 		bytes, err := proto.Marshal(value)
// 		if err != nil {
// 			return errors.Trace(err)
// 		}

// 		taskID := fmt.Sprintf("task:%s", ksuid.New().String())
// 		tasks = append(tasks, taskID)
// 		fields[taskID] = string(bytes)
// 	}

// 	pipeline := queue.db.sess.TxPipeline()
// 	pipeline.HMSet(queue.tasksKey, fields)
// 	pipeline.RPush(queue.queueKey, tasks...)
// 	if _, err := pipeline.Exec(); err != nil {
// 		return errors.Trace(err)
// 	}

// 	log.WithFields(log.Fields{
// 		"queue": queue.queueKey,
// 		"tasks": tasks,
// 	}).Info("Tasks queued")

// 	return nil
// }

// func (queue *Queue) BlockingPop(value proto.Message) (string, error) {
// 	taskID, err := queue.db.sess.BRPopLPush(queue.queueKey, queue.processingKey, 0).Result()
// 	if err != nil {
// 		return "", errors.Trace(err)
// 	}

// 	result, err := queue.db.sess.HGet(queue.tasksKey, taskID).Result()
// 	if err != nil {
// 		return "", errors.Trace(err)
// 	}
// 	if err := proto.Unmarshal([]byte(result), value); err != nil {
// 		return "", errors.Trace(err)
// 	}

// 	log.WithFields(log.Fields{
// 		"queue": queue.queueKey,
// 		"task":  taskID,
// 	}).Info("Task received")

// 	return taskID, nil
// }

// func (queue *Queue) ProcessingLen() (int64, error) {
// 	n, err := queue.db.sess.LLen(queue.processingKey).Result()
// 	if err != nil {
// 		return 0, errors.Trace(err)
// 	}

// 	return n, nil
// }

// func (queue *Queue) Len() (int64, error) {
// 	n, err := queue.db.sess.LLen(queue.queueKey).Result()
// 	if err != nil {
// 		return 0, errors.Trace(err)
// 	}

// 	return n, nil
// }

// type cleanupTask struct {
// 	seen time.Time
// }

// func (queue *Queue) CleanUpProcess(sentryDSN string) {
// 	client := sentry.NewClient(sentryDSN)

// 	seenTasks := map[string]*cleanupTask{}
// 	for {
// 		var err error
// 		seenTasks, err = queue.cleanUp(seenTasks)
// 		if err != nil {
// 			log.WithFields(log.Fields{"error": err.Error(), "stack": errors.ErrorStack(err)}).Error("Clean up failed")
// 			client.ReportInternal(context.Background(), err)
// 		}

// 		time.Sleep(10 * time.Minute)
// 	}
// }

// func (queue *Queue) cleanUp(seenTasks map[string]*cleanupTask) (map[string]*cleanupTask, error) {
// 	tasks, err := queue.db.sess.LRange(queue.processingKey, 0, -1).Result()
// 	if err != nil {
// 		return nil, errors.Trace(err)
// 	}

// 	t := map[string]*cleanupTask{}
// 	for _, task := range tasks {
// 		if s, ok := seenTasks[task]; ok {
// 			if time.Now().Sub(s.seen) > 60*time.Second {
// 				log.WithFields(log.Fields{
// 					"queue": queue.queueKey,
// 					"task":  task,
// 					"seen":  s.seen.String(),
// 				}).Warning("Restoring task that was too long in the processing list")

// 				pipeline := queue.db.sess.TxPipeline()
// 				pipeline.RPush(queue.queueKey, task)
// 				pipeline.LRem(queue.processingKey, 1, task)
// 				if _, err := pipeline.Exec(); err != nil {
// 					return nil, errors.Trace(err)
// 				}
// 			} else {
// 				t[task] = s
// 			}
// 		} else {
// 			t[task] = &cleanupTask{
// 				seen: time.Now(),
// 			}
// 		}
// 	}

// 	return t, nil
// }

// func (queue *Queue) Ack(taskID string) error {
// 	pipeline := queue.db.sess.TxPipeline()
// 	pipeline.HDel(queue.tasksKey, taskID)
// 	pipeline.LRem(queue.processingKey, 1, taskID)
// 	if _, err := pipeline.Exec(); err != nil {
// 		return errors.Trace(err)
// 	}

// 	log.WithFields(log.Fields{
// 		"queue": queue.queueKey,
// 		"task":  taskID,
// 	}).Info("ACK task")
// 	return nil
// }
