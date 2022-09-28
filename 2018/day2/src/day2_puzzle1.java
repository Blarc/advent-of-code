import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.Arrays;

public class day2_puzzle1 {
	public static void main(String[] args) throws IOException {
		
		int wordLen = 26;
		
		if (args.length < 1) {
			System.out.println("manjka vhod");
			System.exit(1);
		}
		
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		
		int res = 0;
		
		char[] token = new char[26];
		
		int sum = 0;
		
		int a = 0;
		int b = 0;
		String readLine;
		while ((readLine = br.readLine()) != null) {
			int[] abc = new int[26];
			token = readLine.toCharArray();
			for (int i = 0; i < token.length; i++) {
				int num = token[i] - 'a';
				abc[num]++;
			}
			
			Boolean bool1 = true;
			Boolean bool2 = true;
			
			for (int i = 0; i < abc.length; i++) {
				
				if (abc[i] == 2 && bool1) {
					bool1 = false;
					a++;
				}
				else if (abc[i] == 3 && bool2) {
					bool2 = false;
					b++;
				}
			}
			
			System.out.println(Arrays.toString(abc) + " " + a +" " + b);
			
		}
		System.out.println(a*b);
		
		
		
	}
}