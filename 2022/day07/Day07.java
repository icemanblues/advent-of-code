import java.io.*;
import java.util.*;

class Day07 {
    public static class File {
        private String name;
        private int size;
        private boolean isFile;
        private File parent;
        private Map<String, File> children;

        public File(String name, int size, boolean isFile, File parent) {
            this.name = name;
            this.size = size;
            this.isFile = isFile;
            this.parent = parent;
            this.children = new HashMap<>();
        }

        public void addFile(File f) {
            children.put(f.name, f);
        }
    }

    public static void sizeOf(File file, Map<File, Integer> memo) {
        if(memo.containsKey(file)) {
            return;
        }

        if(file.isFile) {
            memo.put(file, file.size);
            return;
        }

        int sum = 0;
        for(File f: file.children.values()) {
            sizeOf(f, memo);
            sum += memo.get(f);
        }
        memo.put(file, sum);
    }

    public static void main(String[] args) throws Exception {
        System.out.println("2022 day 07");

        File root = new File("/", 0, false, null);
        File curr = root;

        try(BufferedReader br = new BufferedReader(new FileReader("test.txt"))) {
            for(String line = br.readLine(); line != null; line = br.readLine()) {
                String[] fields = line.split("\s");
                if("$".equals(fields[0])) {
                    String cmd = fields[1];
                    if("cd".equals(cmd)) {
                        switch(fields[2]) {
                            case "..":
                                curr = curr.parent;
                                break;
                            case "/":
                                curr = root;
                                break;
                            default:
                                curr = curr.children.get(fields[2]);
                        }

                    }
                    continue;
                }

                // must be list output
                if("dir".equals(fields[0])) {
                    File dir = new File(fields[1], 0, false, curr);
                    curr.addFile(dir);
                } else {
                    File file = new File(fields[1], Integer.parseInt(fields[0]), true, curr);
                    curr.addFile(file);
                }
            }
        }

        // use memoization to aggregate the size
        Map<File, Integer> memo = new HashMap<>();
        sizeOf(root, memo);
        
        // part 1
        int sum = 0;
        for(File f: memo.keySet()) {
            int s = memo.get(f);
            if(!f.isFile && s <= 100000) {
                sum += memo.get(f);
            }
        }
        System.out.println("Part 1: "+sum);

        int unused = 70000000 - memo.get(root);
        int target = 30000000 - unused;
        int min = memo.get(root);
        for(File f: memo.keySet()) {
            int s = memo.get(f);
            if(!f.isFile && s >= target && s < min) {
                min = s;
            }
        }
        System.out.println("Part 2: "+min);
    }
}
