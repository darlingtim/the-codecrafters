# THE BASE CONVERTER
  #### Concept: Number Systems & strconv (Phase 2 of the CodeCrafterThon Challenge)
  A CLI tool that converts numbers between bases and runs from the terminal.

 ## Rules:  
 1.  → DO NOT USE GOOGLE OR ANY AI TOOL; DO IT YOURSELF - you can share ideas with others, but not copy code.  
 2. → Write everything in Go. Standard library only.  
 3. → Must compile and run without errors.
 4. → Push to your the-codecrafters repo in a folder named: thecodecrafterthon-day2/  
 5. → Include a README.md explaining how to run it

  go-reloaded needs to convert hex and binary strings into decimal numbers. 
  This project teaches you exactly that — and makes you think about what happens when the input is garbage.  

```Text
  Build a CLI tool that converts numbers between
  bases. It runs from the terminal like this:
```
```bash
     > convert 1E hex
       ✦ Decimal: 30

     > convert 10 bin
       ✦ Decimal: 2

     > convert 255 dec
       ✦ Binary:  11111111
       ✦ Hex:     FF
```

  ## Requirements:

  - → Support three input bases: hex, bin, dec.
  - → For dec input, output both binary and hex.
  - → For hex and bin input, output only decimal.
  - → All hex output must be uppercase.
  - → The program keeps running until: quit

  ## Validation — handle all of these cleanly:
  - → "1G" is not valid hex.
  - → "10201" is not valid binary.
  - → "abc" is not a valid decimal.
  - → Negative numbers must be supported for dec.
  - → Empty input must not crash the program.

  ## Think about:
  1. → Which strconv functions handle base conversions for you?  
  2. → How do you detect which characters are valid for a given base?  
  3. → What is the difference between ParseInt and ParseUint — and does it matter here?  


