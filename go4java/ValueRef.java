public class ValueRef {
    // START_INTEGER OMIT
    public static class Integer {
        private int value;

        public Integer(int x) {
            this.value = x;
        }
    }

    // START_REF OMIT
    static void doubleInteger(Integer x) {
        x.value *= 2;
    }

    static void byReference() {
        Integer v = new Integer(2);   // v.value is 2
        doubleInteger(v);             // v.value is 4
    }

    // START_VALUE OMIT
    static void doubleInt(int x) {
        x *= 2;
    }

    static void byValue() {
        int v = 2;     // v is 2
        doubleInt(v);  // v is still 2
    }

    // START_REPLACE OMIT
    static void replaceMyInt(Integer v) {
        v = new Integer(42);
    }

    static void byReferenceIsDifferent() {
        Integer v = new Integer(2);    // v.value is 2
        replaceMyInt(v);               // v.value is still 2!
    }

    // START_MAIN OMIT
    public static void main(String[] args) {
        byReference();
        byValue();
        byReferenceIsDifferent();
    }
}