import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;

public class day1_puzzle1 {
	
	public static void main(String[] args) throws IOException {
		
		if (args.length < 1) {
			System.out.println("manjka vhod");
			System.exit(1);
		}
		
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		
		int res = 0;
		
		String readLine;
		while ((readLine = br.readLine()) != null) {
			res += Integer.parseInt(readLine);
		}
		
		System.out.println(res);
	}
	
}
