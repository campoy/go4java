import java.util.Collection;
import java.util.ArrayList;

class BadInheritance {
    // START_TASK OMIT
    class Task {
        private String message;

        public Task(String message) {
            this.message = message;
        }

        public void Run() {
            System.out.println("running " + this.message);
        }
    }
    // END_TASK OMIT

    // START_RUNNER OMIT
    class Runner {
        private String name;

        public Runner(String name) {
            this.name = name;
        }

        public String getName() {
            return this.name;
        }

        public void Run(Task task) { // HL
            task.Run();
        }

        public void RunAll(Task[] tasks) { // HL
            for (Task task : tasks) {
                Run(task);
            }
        }
    }
    // END_RUNNER OMIT

    // START_COUNTING OMIT
    class CountingRunner extends Runner {
        private int count;

        public CountingRunner(String message) {
            super(message);
            this.count = 0;
        }

        @Override public void Run(Task task) {
            count++; // HL
            super.Run(task);
        }

        @Override public void RunAll(Task[] tasks) {
            count += tasks.length; // HL
            super.RunAll(tasks);
        }

        public int getCount() {
            return count;
        }
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
        new BadInheritance().test();
    }
}