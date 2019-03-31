import java.util.*;
import java.io.*;

public class Day4 {

    public static int countValid(String filename) throws IOException {
	int countValid = 0;

	try(BufferedReader br = new BufferedReader(new FileReader(filename))) {
	    String fileLine = br.readLine();
	    while(fileLine != null) {
		if(isValidNoAnagram(fileLine)) {
		    countValid++;
		}
		
		fileLine = br.readLine();
	    }
	}
	
	
	return countValid;
    }

    public static boolean isPassphraseValid(String passphrase) {
	String[] words = passphrase.split("\\s+");
	Set<String> wordSet = new HashSet<>();
	for(String s: words) {
	    wordSet.add(s);
	}

	return words.length == wordSet.size();
    }
    
    public static boolean isValidNoAnagram(String passphrase) {
	String[] words = passphrase.split("\\s+");
	for(int i=0; i<words.length-1; i++) {
	    for(int j=i+1; j < words.length; j++) {
		if(isAnagram(words[i], words[j]) ) {
		    return false; 
		}
	    }
	}

	return true;
    }

    public static boolean isAnagram(String w1, String w2) {
	if(w1.length() != w2.length()) {
	    return false;
	}

	char[] c1 = w1.toCharArray();
	Arrays.sort(c1);
	char[] c2 = w2.toCharArray();
	Arrays.sort(c2);

	return Arrays.equals(c1, c2);
    }
    
    public static void part1() throws Exception {
	boolean test1 = isPassphraseValid("aa bb cc dd ee");
	System.out.println("true: "+ test1);
	
	boolean test2 = isPassphraseValid("aa bb cc dd aa");
	System.out.println("false: " + test2);
	
       	boolean test3 = isPassphraseValid("aa bb cc dd aaa");
	System.out.println("test: " + test3);

	int c =	countValid("test1.txt");
	System.out.println("should be 2: "+c);

	int answer = countValid("input.txt");
	System.out.println("answer: " + answer);
    }

    public static void part2() throws Exception {
	boolean test1 = isAnagram("cab", "bac");
	System.out.println("true: "+test1);
	
	boolean test2 = isAnagram("cab", "baa");
	System.out.println("false: "+test2);

	boolean test3 = isAnagram("aaa", "aaaa");
	System.out.println("false: "+test3);

	boolean test4 = isValidNoAnagram("abcde fghij");
	System.out.println("true: "+test4);
	
	boolean test5 = isValidNoAnagram("abcde xyz ecdab");
	System.out.println("false: "+test5);
	
	boolean test6 = isValidNoAnagram("a ab abc abd abf abj");
	System.out.println("true: "+test6);
	
	boolean test7 = isValidNoAnagram("iiii oiii ooii oooi oooo");
	System.out.println("true: "+test7);

	boolean test8 = isValidNoAnagram("oiii ioii iioi iiio");
	System.out.println("false: "+test8);

	int answer = countValid("input.txt");
	System.out.println("answer: "+answer);
    }
    
    public static void main(String[] args) throws Exception {
	System.out.println("Day 4");
	
	//	part1();
	part2();
    }
}
