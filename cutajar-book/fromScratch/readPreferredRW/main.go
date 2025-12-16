package fromScratch

import "sync"

type ReadPreferredRWMutex struct {
	readersCounter int
	readersLock    sync.Mutex
	globalLock     sync.Mutex
}

func (rw *ReadPreferredRWMutex) ReadLock() {
	rw.readersLock.Lock()
	rw.readersCounter++
	if rw.readersCounter == 1 {
		rw.globalLock.Lock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadPreferredRWMutex) ReadUnlock() {
	rw.readersLock.Lock()
	rw.readersCounter--
	if rw.readersCounter == 0 {
		rw.globalLock.Unlock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadPreferredRWMutex) WriteLock() {
	rw.globalLock.Lock()
}

func (rw *ReadPreferredRWMutex) WriteUnlock() {
	rw.globalLock.Unlock()
}

/*
This implementation of the readers–writer lock is read-preferring. This means that if there are readers currently holding the lock, a writer trying to acquire the lock will be blocked until all readers have released it. However, if new readers arrive while a writer is waiting, they will be allowed to acquire the lock before the writer, potentially leading to writer starvation.
*/
