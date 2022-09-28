import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.Iterator;
import java.util.LinkedList;

public class day7_puzzle1 {
	
	public static void main(String[] args) throws IOException {
		
		if (args.length < 1) {
			System.out.println("manjka vhod");
			System.exit(1);
		}
		
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		
		byte[] locked = new byte[26];
		LinkedList<Byte>[] steps = new LinkedList[26];
		
		String readLine;
		while ((readLine = br.readLine()) != null) {
			byte cond = (byte)(readLine.charAt(5) - 'A');
			byte step = (byte)(readLine.charAt(36) - 'A');
			
			if (steps[cond] == null) {
				 steps[cond] = new LinkedList<Byte>();
			}
			
			steps[cond].add(step);
			locked[step]++;
		}
		
		for (int i = 0; i < 26; i++) {
			for (int j = 0; j < 26; j++) {
				if (locked[j] == 0) {
					locked[j] = -2;
					System.out.print((char)(j+'A'));
					
					if (steps[j] != null) {
						Iterator<Byte> it = steps[j].iterator();
						while (it.hasNext()) {
							locked[it.next()]--;
						}
						break;
					}
				}
			}
		}
	}

}
