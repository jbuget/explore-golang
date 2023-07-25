# Jour 3 - Web framework

Je voulais faire un endpoint `POST`. Le routeur natif de "http" oblige de gérer à la main, avec des `if-else`. Idem pour la gestion des URLs, *endpoints dynamiques* ou *query params*. J'ai autre chose à faire que parser des string d'URL. Je décdide de passer par un framework Web.

Je le veux :
- léger
- sécurisé
- performant
- au plus proche de la logique de Golant

J'ai hésité toute la journée.

J'ai interrogé Masto et ~~Twitter~~ X.com.

Finalement je suis parti sur **Chi** : 
- plus léger que les autres (Gin, Echo)
- très performant
- plus proche des idiomes de code de Go
- API compatible module natif "http"
- bas niveau (bien pour apprendre)
- Gitidea a migré vers Chi

Autres frameworks envisagés :
- **Fiber** :
  - mais n'implémente pas complètement la norme HTTP (pour des raisons de légèreté et perf) (car basé sur FastHTTP)
  - semble plus complexe à appréhender et le code produit semble aussi plus complexe, moins simple à lire ou faire évoluer
- **Gin** :
  - j'ai énormément hésité car il s'agit du plus gros framework
  - pourrait satisfaire à des besoins de perf (poil moins que Chi)
  - pas mal de recommandation sur Twitter
  - tuto sur le site officiel de Golang
  - semble quand même limité dans la gestion de certains protocoles (dont je n'aurais pas l'usage) tel que Protobuff
  - propose pas mal de middlewares sur étagère → le Express de Golang
- **Echo** :
  - *de bons echos* (lol)
  - autant d'étoiles que Chi
  - j'ai l'impression que c'est à mi-chemin entre Gin et Chi → autant partir sur un framework un peu plus tranché
- **Iris** :
  - semble dépassé

Liens :
- "[Choosing the Right Go Web Framework](https://brunoscheufler.com/blog/2019-04-26-choosing-the-right-go-web-framework)" (2019)
- 
