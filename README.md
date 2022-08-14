# genetics: Heredity and gene mixing

This package provides 64 bit gene sequences as well as the ability to manipulate and mix them for various simulation or procedural generation purposes.

It is nothing fancy, but it does the trick. I have set up an example that encodes some human-ish properties.

## Example gene layout

  _______________________ 2 gender
 || _____________________ 2 eye color
 ||||  __________________ 3 hair color         ___________________________ 4 Openness
 |||| ||| _______________ 4 complexion        ||||  ______________________ 4 Conscientiousness
 |||| |||| ||| __________ 3 height            |||| |||| __________________ 4 Extraversion
 |||| |||| |||| || ______ 3 mass              |||| |||| ||||  ____________ 4 Agreeableness
 |||| |||| |||| |||| | __ 3 growth            |||| |||| |||| ||||  _______ 4 Neuroticism
 |||| |||| |||| |||| ||||                     |||| |||| |||| |||| ||||
 xxxx xxxx|xxxx xxxx|xxxx xxxx|xxxx xxxx|xxxx xxxx|xxxx xxxx|xxxx xxxx|xxxx xxxx
                          |||| |||| |||| ||||                          |||| ||||
 4 strength _________________  |||| |||| ||||                           ________ unused
 4 intelligence __________________  |||| ||||
 4 dexterity __________________________  ||||
 4 resilience ______________________________