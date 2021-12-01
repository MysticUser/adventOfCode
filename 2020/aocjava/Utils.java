package aocjava;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Utils {

    public static List<String> promptInputRows(){
        System.out.println("Paste the input-code here: (send a blank line to continue)");
        Scanner in = new Scanner(System.in);

        List<String> input = new ArrayList<>();

        while(in.hasNextLine()){
            String nextLine = in.nextLine();
            if(nextLine.equals("")){
                break;
            }
            input.add(nextLine);
        }

        return input;
    }

    public static List<Integer> promptInputNumbers(){
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
