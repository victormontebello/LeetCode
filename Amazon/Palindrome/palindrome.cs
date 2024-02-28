namespace Amazon.Palindrome
{
    public class Palindrome
    {
        public bool IsPalindrome(int x)
        {
            if (x < 0)
            {
                return false;
            }

            ReadOnlySpan<char> str = x.ToString();
            for (int i = 0; i < str.Length / 2; i++)
            {
                if (str[i] != str[^(i + 1)])
                {
                    return false;
                }
            }

            return true;
        }
    }
}