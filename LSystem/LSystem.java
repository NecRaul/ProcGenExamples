import java.util.Arrays;
import java.util.List;

public class LSystem {

    private static final String AXIOM = "RNP+BPM;GAdRk2CASeITHwCSIl0THC2AFA5SRFAT=RPMGAF";
    private static final List<String> RULES = Arrays.asList(
            "B => Ca",
            "C => STh",
            "F => Pg",
            "H => OB",
            "M => F",
            "N => nAr",
            "O => CRYF",
            "P => aT",
            "H => TK",
            "T => BP",
            "E => OMg");

    public static void main(String[] args) {
        String newAxiom = AXIOM;
        int iterations = 7;
        for (int i = 0; i <= iterations; i++) {
            System.out.println("Iteration " + (i) + ": " + newAxiom + "\n");
            newAxiom = rewrite(newAxiom);
        }
    }

    private static String rewrite(String newAxiom) {
        StringBuilder builder = new StringBuilder();
        for (char character : newAxiom.toCharArray()) {
            String replacement = findRule(character);
            if (replacement != null) {
                builder.append(replacement);
            } else {
                builder.append(character);
            }
        }
        return builder.toString();
    }

    private static String findRule(char character) {
        for (String rule : RULES) {
            String[] parts = rule.split(" => ");
            String target = parts[0];
            String replace = parts[1];
            if (target.charAt(0) == character) {
                return replace;
            }
        }
        return null;
    }
}
