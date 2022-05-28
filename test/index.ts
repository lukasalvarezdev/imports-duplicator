import { Basic } from './src/basic';
import { DoubleQuotes } from "./src/double-quotes";
import { 
  MultiLine,
  MultiLine2,
} from "./src/multi-line";
import { WithMultipleDots } from "./src/with.multiple.dots";
import { BigFile } from "./src/big.file";

export type { Basic, DoubleQuotes, MultiLine, MultiLine2, WithMultipleDots };
export { BigFile };

// "src/dir0/dir1/dir2/dir3/index.ts"

// "../../dir4/dir5/dir6/import.ts"

// [.., .., dir4, dir5, dir6, import.ts]

// 1st
// "src/dir0/dir1/dir2/dir3/index.ts" -> "src/dir0/dir1/dir2"

// 2nd
// "src/dir0/dir1/dir2/dir3/index.ts" -> "src/dir0/dir1"

// 3rd
// "src/dir0/dir1/dir2/dir3/index.ts" -> "src/dir0/dir1/dir4"

// 4th
// "src/dir0/dir1/dir2/dir3/index.ts" -> "src/dir0/dir1/dir4/dir5"

// 5th
// "src/dir0/dir1/dir2/dir3/index.ts" -> "src/dir0/dir1/dir4/dir5/dir6"

// 6th
// "src/dir0/dir1/dir2/dir3/index.ts" -> "src/dir0/dir1/dir4/dir5/dir6/import.ts"

// if (..) remove last element
// if (.) remove last element only if it is a file
// if (any) add last element