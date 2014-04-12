import java.util.Collection;
import java.util.ArrayList;

class Composition {
    class Task {
        private String message;

        public Task(String message) {
            this.message = message;
        }

        public void Run() {
            System.out.println("running " + this.message);
        }
    }

    class Runner {
        private String name;

        public Runner(String name) {
            this.name = name;
        }

        public String getName() {
            return this.name;
        }

        public void Run(Task task) {
            task.Run();
        }

        public void RunAll(Task[] tasks) {
            for (Task task : tasks) {
                Run(task);
            }
        }
    }

    // START_COUNTING OMIT
    class CountingRunner {
        private Runner runner;
        private int count;

        public CountingRunner(String message) {
            this.runner = new Runner(message);
            this.count = 0;
        }

        public void Run(Task task) {
            count++;
            runner.Run(task);
        }

        public void RunAll(Task[] tasks) {
            count += tasks.length;
            runner.RunAll(tasks);
        }

        public int getCount() { return count; }

        public String getName() { return runner.getName(); }
    }
    // END_COUNTING OMIT

    public void test() {
        // START_MAIN OMIT
        CountingRunner runner = new CountingRunner("my runner");

        Task[] tasks = { new Task("one"), new Task("two"), new Task("three")};
        runner.RunAll(tasks);

        System.out.printf("%s ran %d tasks\n", runner.getName(), runner.getCount());
        // END_MAIN OMIT
    }

    public static void main(String[] args) {
        new Composition().test();
    }
}
