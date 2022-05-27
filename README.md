# Welcome to Imports Duplicator!
A blazingly fast tool that allows you to copy imported files to another directory.

## How to use
Very easy, just download the `exe` file in the root directory and execute it in your favorite terminal.

## Usage

 ```
exe <source_path> <destination_path> <file_to_replace_path>
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
#      This is to replace the file, if you want them in 
#           another file, specify another path
                           ⬇️
exe types.ts temp-types/ types.ts
    
# Your build step
npm run build
```

In the build, `types.ts` would look like this:
```typescript
 import type { User } from './temp-types/user'
 import type { Address } from './temp-types/address'
 
 export type { User, Address }
 ```

## Known limitations
- This tool only copies the files that are referenced in the file that you provide, if those files have imports from another
files, it won't copy them.
- It only works `.ts` files.

## Why would I use this?
This helped my team to solve a  * **really** *  specific case (just temporarily), but it's not intended to have a real use, maybe in the future.

## Extra
I built this project as a part of my learning journey of golang, so you may find this code a bit awful, but I'm open
and would love any suggestions!
