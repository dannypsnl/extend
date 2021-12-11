# CHANGELOG

### Working

### v0.2.0

* ExtBlock
    * prevent overwrite terminator
    * put `phi` instruction to the start of block automatically

### v0.1.0

* NoDupModule
    * wraps all new global definition constructor, rename duplicates
    * `NoDup` removes existed duplicates
* ExtFunc
    * `IsDefintion`
    * `IsDeclaration`
* ExtBlock
    * `HasTerminator`
    * `BelongsToFunc`
