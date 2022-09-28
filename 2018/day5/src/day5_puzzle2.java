import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;


public class day5_puzzle2 {
	
	public static class StackElement
	{
		Object element;
		StackElement next;

		StackElement()
		{
			element = null;
			next = null;
		}
	}

	public static class Stack
	{	
		
		private StackElement top;
		
		public Stack()
		{
			makenull();
		}
		
		public void makenull()
		{
			top = null;
		}
		
		public boolean empty()
		{
			return (top == null);
		}
		
		public Object top()
		{
			if (!empty())
				return top.element;
			else
				return null;
		}
		
		public void push(Object obj)
		{
			StackElement el = new StackElement();
			el.element = obj;
			el.next = top;
			
			top = el;
		}
		
		public void pop()
		{
			if (!empty())
			{
				top = top.next;
			}
		}
	}

	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		
		int c, stackedChar, countRemoved = 0, countAll = -1, max = Integer.MIN_VALUE;
		char[] input = new char[50001];
		
		
		for (int i = 0; (c = br.read()) != -1; i++) {
			countAll++;
			input[i] = (char)c;
		}
		
		
		Stack stack = new Stack();
		
		for (char i = 'A'; i < 'Z'; i++) {
			countRemoved = 0;
			//System.out.println(i + " " + (char)(i+32));
			for (int j = 0; j < input.length; j++) {
				c = input[j];
				if (c == i || c == i+32) {
					countRemoved++;
				}
				
				else if (!stack.empty()) {
					stackedChar = (int)stack.top();
					
					//System.out.println((char)stackedChar + " " + (char)c);
					//System.out.println(stackedChar - c);
					if (stackedChar - c == 32 || stackedChar - c == -32) {
						countRemoved += 2;
						stack.pop();
					}
					else {
						stack.push(c);
					}
				}
				else {
					stack.push(c);
				}
			}
			
			if (countRemoved > max) {
				max = countRemoved;
			}
		}
		
		System.out.println("All: " + countAll);
		System.out.println("Removed: " + max);
		System.out.println("Result: " + (countAll - max));
		
		br.close();
		
	}

}
