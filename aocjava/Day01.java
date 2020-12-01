package aocjava;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

class Day01 {
    public static void main(String[] args) {
        System.out.println("\nAdvent of Code day 1:");

        List<Integer> input = promptInput();

        int partOneSolution = 0;
        int partTwoSolution = 0;

        for (int firstNum : input) {
            for (int secondNum : input) {
                if(firstNum+secondNum == 2020){
                    if(partOneSolution == 0){
                        partOneSolution = firstNum*secondNum;
                    }
                    else if(partOneSolution != firstNum*secondNum){
                        System.out.println("There might be more solutions to part one!");
                    }
                }
                for (int thirdNum : input) {
                    if(firstNum+secondNum+thirdNum == 2020){
                        if(partTwoSolution == 0){
                            partTwoSolution = firstNum*secondNum*thirdNum;
                        }
                        else if(partTwoSolution != firstNum*secondNum*thirdNum){
                            System.out.println("There might be more solutions to part two!");
                        }
                    }
                }
            }
        }

        System.out.println("\nPart one solution: "+partOneSolution);
        System.out.println("Part two solution: "+partTwoSolution);

    }

    private static List<Integer> promptInput(){
        System.out.println("Paste the input-code here, and end with '0' or 'end':");
        Scanner in = new Scanner(System.in);

        List<Integer> input = new ArrayList<>();

        while(in.hasNextLine()){
            if(in.hasNext("e")){
                break;
            }

            int nextNum = in.nextInt();
            if(nextNum == 0){
                break;
            }

            input.add(nextNum);
        }
        return input;
    }
}