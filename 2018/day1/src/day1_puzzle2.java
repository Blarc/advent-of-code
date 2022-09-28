import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;

public class day1_puzzle2 {
	
	public static class Value {
		int val;
		Value next;
		
		public Value (int val) {
			this.val = val;
			this.next = null;
		}
	}
	
	public static class Result {
		int res;
		Result next;
		
		public Result (int res) {
			this.res = res;
			this.next = null;
		}
	}
	
	public static void main(String[] args) throws IOException {
		
		if (args.length < 1) {
			System.out.println("manjka vhod");
			System.exit(1);
		}
		
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		String readLine = br.readLine();
		int nextInt = Integer.parseInt(readLine);
		
		Value firstVal = new Value(nextInt);
		Result firstRes = new Result(nextInt);
		
		Value atmVal = firstVal;
		Result atmRes = firstRes;
		int res = nextInt;
		
		
		while ((readLine = br.readLine()) != null) {
			nextInt = Integer.parseInt(readLine);
			
			res += nextInt;
			
			Result iter = firstRes;
			
			while (iter != null) {
				if (iter.res == res) {
					System.out.println(res);
					break;
				}
				iter = iter.next;
			}
			
			atmVal.next = new Value(nextInt);
			atmRes.next = new Result(res);
			atmVal = atmVal.next;
			atmRes = atmRes.next;
			
			
			
			
		}
		
		/*Value iterVal = firstVal;
		
		while (iterVal != null) {
			System.out.println(iterVal.val);
			iterVal = iterVal.next;
		}*/
		
		/*
		Result iterRes = firstRes;
		
		while (iterRes != null) {
			System.out.println(iterRes.res);
			iterRes = iterRes.next;
		}*/
		
		
		Boolean bool = true;
		Value valIter = firstVal;
		while (bool) {
			res += valIter.val;
			
			Result resIter = firstRes;
			
			while (resIter != null) {
				if (resIter.res == res) {
					System.out.println(res);
					bool = false;
					break;
				}
				resIter = resIter.next;
			}
			
			atmRes.next = new Result(res);
			atmRes = atmRes.next;
			valIter = valIter.next;
			
			if (valIter == null) {
				valIter = firstVal;
			}
		}
		
		
		br.close();
	}
	
}
