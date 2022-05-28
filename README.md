# Welcome to Imports Duplicator!
A blazingly fast tool that allows you to copy imported files to another directory.

## Get started
1. Clone this repository.
2. Set up golang (v1.18)
3. Run `go test`
4. Run `go run . <source_path> <output_file_name>`

If all that worked, you should have now a directory `/out` with all the files that are being referenced in `<source_path>`
and a file `<output_file_name>.ts` with all the info in `<source_path>` but with the imports updated.

That's it!, now it's time to rock!

## How to use
Very easy, just run:
```
go build .
```

Now you should have a binary file `imports-duplicator` in the root directory.

## Usage

This will execute the binary file, you can use your favorite terminal
 ```
imports-duplicator <source_path> <output_file_name>
 ```
 
 ## Example
 
 Your folder structure looks like this:
 
- `components/...`
- `types.ts`
 
The file `types.ts` that looks like this:
 ```typescript
 import type { User } from '../dir_outside_root/user'
 import type { Address } from '../dir_outside_root/address'
 
 export type { User, Address }
 ```
 
 And your code works just fine in dev, but in the build step for some reason you don't have access
 to `../dir_outside_root`. So to have `dir_outside_root/user.ts` 
 and `dir_outside_root/address.ts` avaliable in the build step you would do it like this:
 
 
```sh
go build .
imports-duplicator types.ts import-from-here
    
# Your build step
npm run build
```

In the build, `import-from-here.ts` would look like this:
```typescript
 import type { User } from './out/user'
 import type { Address } from './out/address'
 
 export type { User, Address }
 ```
 
 And your folder structure would look like this:
 
- `components/...`
- `types.ts`
- `import-from-here.ts`
- `out/user.ts`
- `out/address.ts`

## Known limitations
- This tool only copies the files that are referenced in the file that you provide, if those files have imports from another
files, it won't copy them.
- It only works `.ts` files.

## Why would I use this?
This helped my team to solve a  * **really** *  specific case (just temporarily), but it's not intended to have a real use, maybe in the future.

## Extra
I built this project as a part of my learning journey of golang, so you may find this code a bit awful, but I'm open
and would love any suggestions!
