import java.io.PrintStream;

class InterfaceSat {
    // START_PRINT_TO OMIT
    interface StringPrinter {
        void print(String s);
    }

    public void printHiTo(StringPrinter sp) {
        sp.print("hi");
    }
    // END_PRINT_TO OMIT

    /* START_BAD OMIT
        public static void main(String[] args) {
            printHiTo(System.out);
        }
    END_BAD OMIT
    */

    // START_MyStringPrinter OMIT
    class MyStringPrinter implements StringPrinter {
        PrintStream ps;
        MyStringPrinter(PrintStream ps) {
            this.ps = ps;
        }
        public void print(String s) {
            this.ps.println(s);
        }
    }
    // END_MyStringPrinter OMIT

    public void test() {
        printHiTo(new MyStringPrinter(System.out));

        printHiTo(new StringPrinter() {
            public void print(String s) {
                System.out.println(s);
            }
        });
    }

    public static void main(String[] args) {
        new InterfaceSat().test();
    }
}