═══════════════════════════════════════════  
SQUAD PIPELINE CONTRACT  
Squad: [Debuggers]  
───────────────────────────────────────────  
 ## Input line types:
  1. Lines in ALL CAPS
  2. Lines in all lowercase
  3. Lines starting with TODO:
  4. Lines that are only dashes or blanks


 ## Transformation rules (in order):
   1. [Replace TODO: with ✦ ACTION:]
   2. [Convert ALL CAPS lines to Title Case]
   3. [Convert all lowercase lines to uppercase]
   4. [Remove lines that are only dashes or blanks]
   5. [Flag any line longer than 80 characters with [TRUNCATED] at the end]

## Output format:
  ``` Text
   [Header: PROCESSED SENTINEL FIELD REPORT]
   [Line numbering format: "1."]
   [Summary block:  Lines read : [number] ✦ Lines written : [number] ✦ Lines removed : [number] ✦ Rules applied : [] ]

 Terminal summary fields:
   [List what your squad agreed on]
 ═══════════════════════════════════════════

 CodeCrafters — Operation Gopher Protocol
 Module: File Pipeline
 Author: [Timothy Ododo]
 Squad:  [Debuggers]
 ```