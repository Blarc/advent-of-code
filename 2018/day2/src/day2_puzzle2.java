import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.Arrays;

public class day2_puzzle2 {
	public static void main(String[] args) throws IOException {
		
		int wordLen = 26;
		
		if (args.length < 1) {
			System.out.println("manjka vhod");
			System.exit(1);
		}
		
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		
		char[] token = new char[26];
		char[][] all = new char[250][26];
		
		int count = 0;
		String readLine;
		while ((readLine = br.readLine()) != null) {
			all[count] = readLine.toCharArray();
			count++;		
		}
		
		for (int i = 0; i < all.length; i++) {
			for (int j = i+1; j < all.length; j++) {
				int n = 0;
				int index = 0;
				for (int k = 0; k < all[j].length; k++) {				
					if (all[i][k] != all[j][k]) {
						n++;
						index = k;
					}
					if (n > 1) {
						break;
					}
				}
				if (n == 1) {
					all[i][index] = '0';
					System.out.println(all[i]);
					//System.out.println(index);
				}
			}
		}
		
		
	}
}