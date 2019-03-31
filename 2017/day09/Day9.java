import java.io.*;
import java.util.*;

public class Day9 {

    private static class Answer {
	public Answer() {}
	public Answer(int score, int garbageCount) { this.score = score; this.garbageCount = garbageCount; }
	int score = 0;
	int garbageCount = 0;

	public String toString() {
	    return "[score:" + score + " garbageCount:" + garbageCount +"]";
	}
    }
    
    public static Answer score(String s) {
	int score = 0;
	int garbageCount = 0;
	
	int depth = 0;
	boolean inGarbage =  false;
	boolean negateNext = false;
	
	for(int i=0; i<s.length(); i++) {
	    char c = s.charAt(i);

	    if(negateNext) {
		negateNext = false;
		continue;
	    }
	    
	    switch(c) {
	    case '{':
		if(!inGarbage) {
		    depth++;
		} else {
		    garbageCount++;
		}
		break;
	    case '}':
		if(!inGarbage) {
		    score += depth;
		    depth--;
		} else {
		    garbageCount++;
		}
		break;
	    case '!':
		negateNext = true;
		break;
	    case '<':
		if(!inGarbage) {
		    inGarbage = true;
		} else {
		    garbageCount++;
		}
		break;
	    case '>':
		inGarbage = false;
		break;
	    default:
		if(inGarbage) {
		    garbageCount++;
		}
		// no-op
		break;
	    }
		
	}
	
	return new Answer(score, garbageCount);
    }
    
    public static void part1() {
	System.out.println("{}:1 "+ score("{}"));
	System.out.println("{{{}}}:6 "+ score("{{{}}}"));

	System.out.println("{{},{}}:5 "+ score("{{},{}}"));
	System.out.println("{{{},{},{{}}}}:16 "+ score("{{{},{},{{}}}}"));
	System.out.println("{<a>,<a>,<a>,<a>}:1 "+ score("{<a>,<a>,<a>,<a>}"));
	System.out.println("{{<ab>},{<ab>},{<ab>},{<ab>}}:9 "+ score("{{<ab>},{<ab>},{<ab>},{<ab>}}"));
	System.out.println("{{<!!>},{<!!>},{<!!>},{<!!>}}:9 "+ score("{{<!!>},{<!!>},{<!!>},{<!!>}}"));
	System.out.println("{{<a!>},{<a!>},{<a!>},{<ab>}}:3 "+ score("{{<a!>},{<a!>},{<a!>},{<ab>}}"));
	
    }

    public static void part2() {
	System.out.println("<>:0 "+score("<>"));
	System.out.println("<random characters>:17 "+score("<random characters>"));
	System.out.println("<<<<>:3 "+score("<<<<>"));
	System.out.println("<{!>}>:2 "+score("<{!>}>"));

	System.out.println("<!!>:0 "+score("<!!>"));
	System.out.println("<!!!>>:0 "+score("<!!!>>"));
	System.out.println("<{o\"i!a,<{i<a>:10 "+score("<{o\"i!a,<{i<a>"));
    }
    
    public static void main(String[] args) throws Exception {
	System.out.println("Day9 started!");

	part1();
	try(BufferedReader br = new BufferedReader(new FileReader("input.txt"))) {
	    String question = br.readLine();
	    Answer answer = score(question);
	    System.out.println("answer: "+answer);
	}
	

	part2();
    }
}
