# Jour 6 - API et concurrence

## Fetching API

Aujourd'hui, j'ai voulu voir comment consommer des API tierces / externes.

Naïvement, j'ai cherché à reproduire ce que je connais en Node.js, c'est-à-dire le fait d'utiliser le pattern `async` / `await`, pour chaque traitement asynchrone.

J'ai passé (perdu) pas mal de temps à lire de la doc et des articles.
Il y a pas mal de littérature sur le sujet.
J'ai compris le principe : il faut mixer `goroutines` (traitement non-bloquant qui s'exécute dans un thread et qu'on peut cumuler dans ce même thread), `channels` (tuyaux qui permettent de partager de la mémoire / info entre plusieurs goroutines) et les `waitGroups` (repdoruction peu ou prou de `Promise.all()`).

Ça me paraissait commpliqué et "j'intuitais" que ce n'était pas la façon de faire en Go.
J'ai réorienté mes recherches non plus sur "golang async await" mais plutôt sur "golang fetch api GET POST".
Bien m'en a pris.
Je suis tombé sur un autre type de docu / articles.

J'ai testé (je n'avais pas encore pondu le moindre code jusqu'alors).
Ça a fonctionné (cf. commit #bc48194).

> 💡 J'ai utilisé l'API PokdeAPI : https://pokeapi.co/api/v2/pokemon/charizard

Liens :
- https://betterprogramming.pub/how-to-use-async-await-in-go-b595f950aa6
- https://go.dev/tour/concurrency/2
- https://go.dev/tour/concurrency/1
- https://hackernoon.com/asyncawait-in-golang-an-introductory-guide-ol1e34sg
- ⭐️ https://madeddu.xyz/posts/go-async-await/ 

## Premiers endpoints

J'ai commencé à préparer le terrain et à implémenter un premier endpoint facile : `GET /accounts/me`.

Je ne maîtrise vraiemnt pas suffisamment la techno pour y aller en TDD dès le premier endpoint.

J'obtiens un premier résultat satisfaisant.

Par contre, comme j'utilise le minimum de dépendances, ça m'oblige à faire BEAUCOUP de choses àla main, ce qui est un peu fastidieux.
Mais c'est comme ça qu'on apprend 🤷‍♂️