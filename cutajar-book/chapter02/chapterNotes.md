
<details>
<summary>2.1 What is the primary benefit of multiprocessing (multiprogramming) regarding CPU usage?</summary>

It maximizes CPU utilization. When one job idles (e.g., waiting for I/O), the OS switches context to run another job so the processor isn't wasted.

</details>

<details>
<summary>2.1 In the lifecycle of a job, what happens when a running process requests an I/O operation?</summary>

The OS removes the job from the CPU and places it in the  **I/O waiting queue** . The CPU then picks up a different job from the **ready queue** to keep processing.

</details>

<details>
<summary>2.1 What is an Interrupt, and what component handles it?</summary>

An interrupt is a signal to stop current execution and notify the system of an event (like I/O completion or a timer tick). It is managed by the  **interrupt controller** , which tells the CPU to switch tasks.

</details>

<details>
<summary>2.1 What is a Context Switch?</summary>

The process of stopping one job and starting another. The OS saves the current state (registers, program counter) to a **Process Control Block (PCB)** and loads the PCB of the next job.

</details>

<details>
<summary>2.2 What is the fundamental difference in memory handling between a Process and a Thread?</summary>

* **Processes** have isolated memory spaces (secure but resource-heavy).
* **Threads** share the same memory space/heap within a process (efficient/fast but requires synchronization).

</details>

<details>
<summary>2.2 What does the fork() system call do on UNIX systems?</summary>

It creates a child process that is a complete copy of the parent (memory, registers, file handles). It returns `0` to the child and the child's PID to the parent.

</details>

<details>
<summary>2.2 What is Copy on Write (COW) optimization in fork()?</summary>

The OS does not copy the entire memory space immediately. Parent and child share memory pages until one tries to write to them; only then is that specific page copied.

</details>

<details>
<summary>2.2 Although threads share the heap, what memory resource is private to each thread?</summary>

The  **Stack** . Each thread needs its own stack to track local variables, function calls, and the program counter independently.

</details>

<details>
<summary>2.2 Why is switching between threads faster than switching between processes?</summary>

Threads live in the same virtual memory space, so the OS doesn't need to switch out memory mappings or flush distinct caches, requiring fewer resources to manage.

</details>

<details>
<summary>2.3 Why does the Go runtime need to sleep (e.g., time.Sleep) in the main function when launching goroutines in simple examples?</summary>

Because if the `main` function (the main thread) finishes execution, the **entire process terminates** immediately, regardless of whether child goroutines are still running.

</details>

<details>
<summary>2.3 What is the difference between User-Level Threads and Kernel-Level Threads?</summary>

* **Kernel-Level:** Managed by the OS; expensive context switch; if one blocks on I/O, others can still run.
* **User-Level:** Managed by the runtime (user space); fast context switch; but if one makes a blocking system call, the **entire process (and all threads)** blocks.

</details>

<details>
<summary>2.3 Which threading model does Go use?</summary>

The  **M:N Hybrid Model** . It maps **M** goroutines (user-level) onto **N** OS threads (kernel-level).

</details>

<details>
<summary>2.3 How does Go determine the default number of Kernel-Level threads (N) to utilize?</summary>

It uses the `GOMAXPROCS` variable, which defaults to the number of logical CPUs (`runtime.NumCPU()`) available on the hardware.

</details>

<details>
<summary>2.3 How does Go handle a goroutine making a blocking system call (like file I/O) without blocking the other goroutines?</summary>

Go detaches the blocked goroutine and its underlying OS thread from the processor. It then spins up a **new (or cached) OS thread** to continue executing the remaining goroutines in the local run queue.

</details>

<details>
<summary>2.3 What is "Work Stealing" in the Go scheduler?</summary>

If a kernel-level thread runs out of goroutines in its  **Local Run Queue (LRQ)** , it attempts to "steal" half the goroutines from another thread's queue to balance the workload.

</details>

<details>
<summary>2.3 Why is Go's scheduler considered "Cooperative" (at the user level) compared to the OS "Preemptive" scheduler?</summary>

The OS interrupts threads via hardware timers (preemptive). The Go scheduler relies on **user-level events** (function calls, channel ops, `runtime.Gosched()`) to trigger a context switch; it cannot arbitrarily pause a running instruction stream without these triggers.

</details>

<details>
<summary>2.4 Succinctly define the difference between Concurrency and Parallelism.</summary>

* **Concurrency:** The *structure* of a program that handles multiple tasks at once (dealing with lots of things).
* **Parallelism:** The *execution* of multiple tasks exactly simultaneously (doing lots of things at once).

</details>

<details>
<summary>2.4 Can a single-core processor exhibit parallelism?</summary>

No. It can only exhibit **concurrency** (via time-slicing/interleaving). True parallelism requires multiple physical processing units executing instructions at the exact same instant.

</details>

<details>
<summary>2.4 What is Pseudo-Parallelism?</summary>

When a single-processor system switches context so rapidly that it gives the **illusion** that jobs are running in parallel, though they are actually running sequentially in slices.

</details>
