import java.io.*;
import java.util.*;

public class Day8 {
    private static final String MAX = "8";
    
    private static int processInstructions(String filename) throws IOException {
	Map<String, Integer> registers = new HashMap<>();
	
	try(BufferedReader br = new BufferedReader(new FileReader(filename))) {
	    String fileLine = br.readLine();
	    while(fileLine != null) {
		process(fileLine, registers);
		fileLine = br.readLine();
	    }
	}

	int max = Integer.MIN_VALUE;
	for(Integer i: registers.values()) {
	    if( i > max) {
		max = i;
	    }
	}

	//return max;
	return registers.getOrDefault(MAX, 0);
    }

    private static void process(String instruction, Map<String, Integer> registers) {
	String[] tokens = instruction.split("\\s+");
	// b inc 5 if a > 1
	// 0   1 2  3 4 5 6
	String predicateVariable = tokens[4];
	int predicateVariableValue = registers.getOrDefault(predicateVariable, 0);
	int predicateOperand = Integer.parseInt(tokens[6]);
	String predicateOperator = tokens[5];

	boolean predicate = false;
	switch(predicateOperator) {
	case ">":
	    predicate = predicateVariableValue > predicateOperand;
	    break;
	case "<":
	    predicate = predicateVariableValue < predicateOperand;
	    break;
	case "==":
	    predicate = predicateVariableValue == predicateOperand;
	    break;
	case ">=":
	    predicate = predicateVariableValue >= predicateOperand;
	    break;
	case "<=":
	    predicate = predicateVariableValue <= predicateOperand;
	    break;
	case "!=":
	    predicate = predicateVariableValue != predicateOperand;
	    break;
	}

	if(predicate) {
	    int value = registers.getOrDefault(tokens[0], 0);
	    int operand = Integer.parseInt(tokens[2]);
	    if("inc".equals(tokens[1])) {
		int newValue = value + operand;
		registers.put(tokens[0], newValue);

		int max = registers.getOrDefault(MAX, 0);
		if(newValue > max) {
		    registers.put(MAX, newValue);
		}
	    }
	    else {
		int newValue = value - operand;
		registers.put(tokens[0], newValue);

		int max = registers.getOrDefault(MAX, 0);
		if(newValue > max) {
		    registers.put(MAX, newValue);
		}
	    }
	}
    }

    public static void part1() throws Exception {
	int max1 = processInstructions("test1.txt");
	System.out.println("max1 = 1 : " + max1);

	int answer = processInstructions("input.txt");
	System.out.println("answer: " + answer);
    }
    
    public static void main(String[] args) throws Exception {
	System.out.println("Day 8");
	part1();
    }
}
