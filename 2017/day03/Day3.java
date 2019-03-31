public class Day3 {

    public static int numSteps(int n) {
        return 0;
    }

    public static int[][] buildMemory(int n) {
        double sqrt = Math.sqrt((double) n);
        int length = (int) Math.ceil(sqrt);

        int[][] memory = new int[length][length];
        return memory;
    }

    public static void part1() {
        // 1    -> 0
        int test1 = numSteps(1);
        System.out.println("1->0 : "+test1);

        // 12   -> 3
        int test2 = numSteps(12);
        System.out.println("12->3 : "+test2);

        // 23   -> 2
        int test3 = numSteps(23);
        System.out.println("23->2 : "+test3);

        // 1024 -> 31
        int test4 = numSteps(1024);
        System.out.println("1024->31 : "+test4);

        // 312051
        int answer = numSteps(312051);
        System.out.println("answer: "+answer);
    }

    public static void main(String[] args) {
        System.out.println("Day 3");

        part1();
    }
}