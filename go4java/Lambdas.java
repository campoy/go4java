class Lambdas {
    // START OMIT
    interface Task {
        void run();
    }

    class Runner {
        void run(Task t) {
            t.run();
        }
    }
    // END OMIT
    void test() {
        // ANONYMOUS OMIT
        Runner runner = new Runner();
        runner.run(new Task() {
            public void run() {
                System.out.println("an anonymous interface");
            };
        });
        // LAMBDA OMIT
        runner.run(() ->System.out.println("a lambda expression"));
    }

    public static void main(String[] args) {
        new Lambdas().test();
    }
}