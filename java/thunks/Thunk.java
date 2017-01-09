package thunks;

import java.io.File;
import java.io.InputStream;
import java.nio.file.Files;
import java.nio.file.StandardCopyOption;

public class Thunk {
    public final static String REAL_EXE_NAME = "bins/exefile";
    public final static String FILE_SUFFIX = System.getProperty("os.name").startsWith("Windows") ? ".exe" : "";

    private String[] mArgs;
    private String mFileName;

    public Thunk(String[] args) {
        mArgs = args;
    }

    public Thunk load() throws Exception {
        InputStream is = getClass().getClassLoader().getResourceAsStream(REAL_EXE_NAME);
        File exefile = File.createTempFile("thunk_exe", FILE_SUFFIX);

        Files.copy(is, exefile.toPath(), StandardCopyOption.REPLACE_EXISTING);

        exefile.setExecutable(true);
        exefile.deleteOnExit();

        mFileName = exefile.getAbsolutePath();

        return this;
    }

    public int execute() throws Exception {
        String[] args = new String[mArgs.length + 1];

        System.arraycopy(mArgs, 0, args, 1, mArgs.length);
        args[0] = mFileName;

        return new ProcessBuilder(args).inheritIO().start().waitFor();
    }

    public static void main(String[] args) throws Exception {
        System.exit(new Thunk(args).load().execute());
    }
}
