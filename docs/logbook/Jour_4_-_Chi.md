# Jour 4 - Chi

Remplacer la config built-in du module "http" par le routeur de Chi fut quasi-immédiat, même pour les endpoints `POST`.
Pour le coup, la promesse est vraiment tenue.

J'ai commencé par faire tourner Chi.

Puis j'ai ajouté les middlewares conseillé dans la doc (RequestID, RealIP, Logger, Recoverer).

Pour les logs, il faudra remplacer les logs en texte par des logs JSON, car c'ets plus pratique pour l'exploitation dans un moteur de logs (Typesense, ELK, Grafana).

Ensuite, j'ai ajouté le middleware Chi-CORS, pour checker vite fait comment ça se passe côté sécu. Facile.
On m'a dit que Gin était mieux pourvu en termes de middlewares, etc.
Pour l'instant, je trouve que c'est fluide et je retrouve le minimum vital attendu avec Chi.

J'ai tenté d'utiliser le middleware `docgen` mais il n'y a rien dans la doc, la lib n'est plus maintenue depuis nov. 2022, et je ne comprends pas simplement comment faire.
Ca m'a saoulé.
J'ai lâché l'affaire.

J'ai commencé à regarder pour avoir une tâche qui *watch* les sources (comme Nodemon pour Node.js).

Les choses que j'aimerais maîtriser : 
- tester
  - https://pkg.go.dev/testing
  - https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/
  - https://github.com/onsi/ginkgo
  - https://elk.zone/piaille.fr/@jbuget/110781459738617643
  - Sur conseil de Gabriel, je décide de partir sur Testify (20K stars) ~ https://github.com/stretchr/testify
    - le module "testing" est trop *brut* à mon goût
    - je préfère une syntaxe et des outils un peu plus haut niveau, façon light-BDD 
- debugger
- source watcher
- deployer
- task manager
  - [taskfile.dev](https://taskfile.dev/)
  - [GNU Make](https://tutorialedge.net/golang/makefiles-for-go-developers/) for Golang devs
  - [autre article](https://earthly.dev/blog/golang-makefile/)
- importer modules internes

Pour le **task manager**, je préfère utiliser `GNU Make` plutôt que `Taskfile` car je n'ai rien eu à installer sur mon Mac.
Juste créer un fichier Makefile et faire `make run`.
Ca me paraît plus "standard".
Sur VS Code, il y a une extension opour Makefile : `Makefile Tools`.

Pour **déboguer**, c'était hyper simple. J'a isuivi [cet article](https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code) mais j'aurais pu y aller à l'instinct / habitude.
Ça a fonctionné du premier coup.


## Misc

Installer Chi :

```
go get -u github.com/go-chi/chi/v5
```

Installer Chi-CORS :

```
go get github.com/go-chi/cors
```

