using System.Text;

string input = "1";
int iterations = 10;
int currentIteration = 0;
char[] characterArray = input.ToCharArray();
do
{
    Console.WriteLine($"Iteration {currentIteration}: {new string(characterArray)}");
    StringBuilder result = new StringBuilder();
    char currentCharacter = characterArray[0];
    int currentLength = 0;
    for (int i = 0; i < characterArray.Length; i++)
    {
        if (currentCharacter == characterArray[i])
        {
            currentLength++;
        }
        else
        {
            result.Append(currentLength);
            result.Append(currentCharacter);
            currentLength = 1;
            currentCharacter = characterArray[i];
        }
    }
    result.Append(currentLength);
    result.Append(currentCharacter);
    characterArray = result.ToString().ToCharArray();
    currentIteration++;
}
while (currentIteration <= iterations);
