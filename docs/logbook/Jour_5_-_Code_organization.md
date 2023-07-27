# Jour 5 - Code organization

Je me suis rendu compte lors du précédent jour, en voulant faire des tests ❌, que créer des *packages* et les appeler au sein d'un *module* n'est pas si simple.

C'est parti pour lire "[Organizing Go code](https://go.dev/blog/organizing-go-code)" !

[Apparemment](https://www.jetbrains.com/help/go/configuring-goroot-and-gopath.html#gopath), il est important de définir les variables d'environnement `GOROOT` et `GOPATH`.

Lisons aussi [GOPATH and Modules](https://pkg.go.dev/cmd/go#hdr-GOPATH_and_Modules).
❌ Illisible (dans un délai que je me suis raisonnablement fixé).

✅ ZE article à lire : "[OK Let’s Go: Three Approaches to Structuring Go Code
](https://www.humansecurity.com/tech-engineering-blog/ok-lets-go-three-approaches-to-structuring-go-code)" (2019)
Présente 3 approches.
Il ne faut pas mixer les approches.
J'opte pour l'approche simpliste #1 : mono-package / flat-hierarchy.

❌ Ca compile dans VS Code mais ça ne fonctionne pas quand je fais `make run` ou même `go run main.go`.
Je change l'arborescence (supprime `/src` et remets tout à la racine).
❌ Pas mieux.

C'est pénible (et inusuel, comme cité dans la doc) cette façon d'organiser les fichiers.

✅ A minima, les tests via la commande `go test` passent cette fois.

Je tente une autre approche.
Je remets le dossier + package `hello`.
Je laisse faire l'autocomplétion de VS Code.
Il me sort un import `"jbuget.fr/explore-golang/hello"`.
La compilation se passe bien.
Je lance le programme `go run main.go`.
✅ Ça fonctionne !
❌ Par contre, il ne détecte à nouveau plus les tests 🤦‍♂️


