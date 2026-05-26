package workerpool

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/gammazero/deque"
)

var ErrStopped = errors.New("submitting work to stopped workerpool")

// New creates and starts a pool of worker goroutines.
//
// The maxWorkers parameter specifies the maximum number of workers that can
// execute tasks concurrently. When there are no incoming tasks, workers are
// gradually stopped until there are no remaining workers.
func New(maxWorkers int, options ...Option) *WorkerPool {
	_ = "STUB: not implemented"
	// There must be at least one worker.
	return nil
}

// Start the task dispatcher.

// WorkerPool is a collection of goroutines, where the number of concurrent
// goroutines processing requests does not exceed the specified maximum.
type WorkerPool struct {
	maxWorkers   int
	taskQueue    chan func()
	workerQueue  chan func()
	stoppedChan  chan struct{}
	stopSignal   chan struct{}
	waitingQueue deque.Deque[func()]
	stopLock     sync.Mutex
	stopOnce     sync.Once
	stopped      bool
	waiting      int32
	wait         bool
}

// Size returns the maximum number of concurrent workers.
func (p *WorkerPool) Size() int { _ = "STUB: not implemented"; return 0 }

// Stop stops the worker pool and waits for only currently running tasks to
// complete. Pending tasks that are not currently running are abandoned. Tasks
// must not be submitted to the worker pool after calling stop.
//
// Since creating the worker pool starts at least one goroutine, for the
// dispatcher, Stop() or StopWait() should be called when the worker pool is no
// longer needed.
func (p *WorkerPool) Stop() {
	_ = "STUB: not implemented"

	// StopWait stops the worker pool and waits for all queued tasks tasks to
	// complete. No additional tasks may be submitted, but all pending tasks are
	// executed by workers before this function returns.
	return
}

func (p *WorkerPool) StopWait() {
	_ = "STUB: not implemented"

	// Stopped returns true if this worker pool has been stopped.
	return
}

func (p *WorkerPool) Stopped() bool { _ = "STUB: not implemented"; return false }

// Do enqueues a function for a worker to execute. Returns ErrStopped if the
// worker pool is stopped.
//
// Any external values needed by the task function must be captured in a
// closure. Any return values should be returned over a channel that is
// captured in the task function closure.
//
// Do will not block regardless of the number of tasks submitted. Each task
// is immediately given to an available worker or to a newly started worker. If
// there are no available workers, and the maximum number of workers are
// already created, then the task is put onto a waiting queue.
//
// When there are tasks on the waiting queue, any additional new tasks are put
// on the waiting queue. Tasks are removed from the waiting queue as workers
// become available.
//
// As long as no new tasks arrive, one available worker is shutdown each time
// period until there are no more idle workers. Since the time to start new
// goroutines is not significant, there is no need to retain idle workers
// indefinitely.
func (p *WorkerPool) Do(task func()) (err error) { _ = "STUB: not implemented"; return nil }

// Submit enqueues the given function without blocking. Panics if the worker
// pool is stopped.
func (p *WorkerPool) Submit(task func()) { _ = "STUB: not implemented"; return }

// SubmitWait enqueues the given function, without blocking, and waits for it
// to be executed. Panics if the worker pool is stopped.
func (p *WorkerPool) SubmitWait(task func()) { _ = "STUB: not implemented"; return }

// WaitingQueueSize returns the count of tasks in the waiting queue.
func (p *WorkerPool) WaitingQueueSize() int { _ = "STUB: not implemented"; return 0 }

// Pause causes all workers to wait on the given Context, thereby making them
// unavailable to run tasks. Pause returns when all workers are waiting. Tasks
// can continue to be queued to the workerpool, but are not executed until the
// Context is canceled or times out.
//
// Calling Pause when the worker pool is already paused causes Pause to wait
// until all previous pauses are canceled. This allows a goroutine to take
// control of pausing and unpausing the pool as soon as other goroutines have
// unpaused it.
//
// When the workerpool is stopped, workers are unpaused and queued tasks are
// executed during StopWait.
func (p *WorkerPool) Pause(ctx context.Context) { _ = "STUB: not implemented"; return }

// Wait for workers to all be paused

// dispatch sends the next queued task to an available worker.
func (p *WorkerPool) dispatch(idleTimeout time.Duration) { _ = "STUB: not implemented"; return }

// As long as tasks are in the waiting queue, incoming tasks are put
// into the waiting queue and tasks to run are taken from the waiting
// queue. Once the waiting queue is empty, then go back to submitting
// incoming tasks directly to available workers.

// Got a task to do.

// Create a new worker, if not at max.

// Enqueue task to be executed by next available worker.

// Timed out waiting for work to arrive. Kill a ready worker if
// pool has been idle for a whole timeout.

// If instructed to wait, then run tasks that are already queued.

// Stop all remaining workers as they become ready.

// worker executes tasks and stops when it receives a nil task.
func worker(task func(), workerQueue chan func(), wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

// stop tells the dispatcher to exit, and whether or not to complete queued
// tasks.
func (p *WorkerPool) stop(wait bool) { _ = "STUB: not implemented"; return }

// Signal that workerpool is stopping, to unpause any paused workers.

// Acquire stopLock to wait for any pause in progress to complete. All
// in-progress pauses will complete because the stopSignal unpauses the
// workers.

// The stopped flag prevents any additional paused workers. This makes
// it safe to close the taskQueue.

// Close task queue and wait for currently running tasks to finish.

// processWaitingQueue puts new tasks onto the waiting queue, and removes
// tasks from the waiting queue as workers become available. Returns false if
// worker pool is stopped.
func (p *WorkerPool) processWaitingQueue() bool { _ = "STUB: not implemented"; return false }

// A worker was ready, so gave task to worker.

func (p *WorkerPool) killIdleWorker() bool { _ = "STUB: not implemented"; return false }

// Sent kill signal to worker.

// No ready workers. All, if any, workers are busy.

// runQueuedTasks removes each task from the waiting queue and gives it to
// workers until queue is empty.
func (p *WorkerPool) runQueuedTasks() { _ = "STUB: not implemented"; return }

// A worker is ready, so give task to worker.
