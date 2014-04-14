public class ValueRef {
    public static class MyInteger {
        private int value;
        public MyInteger(int x) {
            this.value = x;
        }
    }

    static void doubleMyInteger(MyInteger x) {
        x.value *= 2;
    }
    static void byReference() {
        MyInteger v = new MyInteger(2);  // v.value is 2
        doubleMyInteger(v);              // v.value is 4
    }

    static void doubleInt(int x) {
        x *= 2;
    }
    static void byValue() {
        int v = 2;     // v is 2
        doubleInt(v);  // v is still 2
    }

    static void replaceMyInteger(MyInteger v) {
        v = new MyInteger(42);
    }
    static void byReferenceIsDifferent() {
        MyInteger v = new MyInteger(2);  // v.value is 2
        replaceMyInteger(v);             // v.value is still 2!
    }

    // START_MAIN OMIT
    public static void main(String[] args) {
        byReference();
        byValue();
        byReferenceIsDifferent();
    }
}