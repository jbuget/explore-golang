# Jour 7 - True case

Apprentissages
- Générer une string random
- Concaténer 2 strings
- Exécuter un INSERT en BDD et récupérer l'ID (`DB.QueryRow` vs. `DB.Exec`)
- Il n'y a pas d'héritage en Go, seulement composition
- en PG, pas de `DateTime` mais `Timestamp`
- en Go, le type "date du jour avec timestamp" est `time.Time`
- la `valeur zéro` pour un type time.Time est `time.Time{}`
- pour tester la valeur zéro d'une date c'est `time.IsZero()`
- obligé de mettre les attributs public d'un struct en majuscule sinon on ne les voit pas

Décisions
- nommer "Insert" plutôt que "Create" ou "Save" pour un ajout d'une entrée en BDD (dans le *store*)
- la création d'un compte / une entité est vue comme un process métier, donc pas dans un repository, que je décide de voir uniquement comme un machin technique
- je décide quand même que `created_at` aura une valeur `NOW()` par défaut côté SQL pour blinder la cohérence des data
