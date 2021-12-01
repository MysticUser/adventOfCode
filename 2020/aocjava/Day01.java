package aocjava;

import java.util.List;

class Day01 {
    public static void main(String[] args) {
        System.out.println("\nAdvent of Code day 1:");

        List<Integer> input = Utils.promptInputNumbers();

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
}