
<details>
<summary><strong>2.1 What is the primary benefit of multiprocessing (multiprogramming) regarding CPU usage?</strong></summary>

It maximizes CPU utilization. When one job idles (e.g., waiting for I/O), the OS switches context to run another job so the processor isn't wasted.

</details>

<details>
<summary><strong>2.1 In the lifecycle of a job, what happens when a running process requests an I/O operation?</strong></summary>

The OS removes the job from the CPU and places it in the `<strong>`I/O waiting queue `</strong>`. The CPU then picks up a different job from the `<strong>`ready queue `</strong>` to keep processing.

</details>

<details>
<summary><strong>2.1 What is an Interrupt, and what component handles it?</strong></summary>

An interrupt is a signal to stop current execution and notify the system of an event (like I/O completion or a timer tick). It is managed by the `<strong>`interrupt controller `</strong>`, which tells the CPU to switch tasks.

</details>

<details>
<summary><strong>2.1 What is a Context Switch?</strong></summary>

The process of stopping one job and starting another. The OS saves the current state (registers, program counter) to a `<strong>`Process Control Block (PCB)`</strong>` and loads the PCB of the next job.

</details>

<details>
<summary><strong>2.2 What is the fundamental difference in memory handling between a Process and a Thread?</strong></summary>

* `<strong>`Processes `</strong>` have isolated memory spaces (secure but resource-heavy).
* `<strong>`Threads `</strong>` share the same memory space/heap within a process (efficient/fast but requires synchronization).

</details>

<details>
<summary><strong>2.2 What does the <code>fork()</code> system call do on UNIX systems?</strong></summary>

It creates a child process that is a complete copy of the parent (memory, registers, file handles). It returns `<code>`0 `</code>` to the child and the child's PID to the parent.

</details>

<details>
<summary><strong>2.2 What is Copy on Write (COW) optimization in <code>fork()</code>?</strong></summary>

The OS does not copy the entire memory space immediately. Parent and child share memory pages until one tries to write to them; only then is that specific page copied.

</details>

<details>
<summary><strong>2.2 Although threads share the heap, what memory resource is private to each thread?</strong></summary>

The `<strong>`Stack `</strong>`. Each thread needs its own stack to track local variables, function calls, and the program counter independently.

</details>

<details>
<summary><strong>2.2 Why is switching between threads faster than switching between processes?</strong></summary>

Threads live in the same virtual memory space, so the OS doesn't need to switch out memory mappings or flush distinct caches, requiring fewer resources to manage.

</details>

<details>
<summary><strong>2.3 Why does the Go runtime need to sleep (e.g., <code>time.Sleep</code>) in the main function when launching goroutines in simple examples?</strong></summary>

Because if the `main` function (the main thread) finishes execution, the `<strong>`entire process terminates `</strong>` immediately, regardless of whether child goroutines are still running.

</details>

<details>
<summary><strong>2.3 What is the difference between User-Level Threads and Kernel-Level Threads?</strong></summary>

* `<strong>`Kernel-Level:`</strong>` Managed by the OS; expensive context switch; if one blocks on I/O, others can still run.
* `<strong>`User-Level:`</strong>` Managed by the runtime (user space); fast context switch; but if one makes a blocking system call, the `<strong>`entire process (and all threads)`</strong>` blocks.

</details>

<details>
<summary><strong>2.3 Which threading model does Go use?</strong></summary>

The `<strong>`M:N Hybrid Model `</strong>`. It maps `<strong>`M `</strong>` goroutines (user-level) onto `<strong>`N `</strong>` OS threads (kernel-level).

</details>

<details>
<summary><strong>2.3 How does Go determine the default number of Kernel-Level threads (N) to utilize?</strong></summary>

It uses the `GOMAXPROCS` variable, which defaults to the number of logical CPUs (`runtime.NumCPU()`) available on the hardware.

</details>

<details>
<summary><strong>2.3 How does Go handle a goroutine making a blocking system call (like file I/O) without blocking the other goroutines?</strong></summary>

Go detaches the blocked goroutine and its underlying OS thread from the processor. It then spins up a `<strong>`new (or cached) OS thread `</strong>` to continue executing the remaining goroutines in the local run queue.

</details>

<details>
<summary><strong>2.3 What is "Work Stealing" in the Go scheduler?</strong></summary>

If a kernel-level thread runs out of goroutines in its `<strong>`Local Run Queue (LRQ)`</strong>`, it attempts to "steal" half the goroutines from another thread's queue to balance the workload.

</details>

<details>
<summary><strong>2.3 Why is Go's scheduler considered "Cooperative" (at the user level) compared to the OS "Preemptive" scheduler?</strong></summary>

The OS interrupts threads via hardware timers (preemptive). The Go scheduler relies on `<strong>`user-level events `</strong>` (function calls, channel ops, `runtime.Gosched()`) to trigger a context switch; it cannot arbitrarily pause a running instruction stream without these triggers.

</details>

<details>
<summary><strong>2.4 Succinctly define the difference between Concurrency and Parallelism.</strong></summary>

* `<strong>`Concurrency:`</strong>` The `<em>`structure `</em>` of a program that handles multiple tasks at once (dealing with lots of things).
* `<strong>`Parallelism:`</strong>` The `<em>`execution `</em>` of multiple tasks exactly simultaneously (doing lots of things at once).

</details>

<details>
<summary><strong>2.4 Can a single-core processor exhibit parallelism?</strong></summary>

No. It can only exhibit `<strong>`concurrency `</strong>` (via time-slicing/interleaving). True parallelism requires multiple physical processing units executing instructions at the exact same instant.

</details>

<details>
<summary><strong>2.4 What is Pseudo-Parallelism?</strong></summary>

When a single-processor system switches context so rapidly that it gives the `<strong>`illusion `</strong>` that jobs are running in parallel, though they are actually running sequentially in slices.

</details>
