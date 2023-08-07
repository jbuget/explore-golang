# Jour 12 - J'en Chi (non)

Après plusieurs jours de prise en main et d'exploration, je trouve que Chi est un excellent framework Web et fait parfaitement le job que j'attend de ce type de technologie.

Je ne saurais dire quand et pourquoi, mais je ressens que j'ai déjà pu bénéficié et apprécié le côté "**compatible au module built-in `http`**".

Les **fonctionnalités de routing et l'expérience développeur** associée sont exactement celles que j'attends.
Moment cool : j'ai réussi à designer un endpoint dynamique du premier coup sans même avoir besoin d'accéder à la documentation (je n'avais pas Internet à ce moment là) !

Il y a tous les **middlewares** minimum et communs nécessaires à tous projets un peu sérieux / ambitieux :
- sécurité
- perf / compression
- praticité
- authentification / contextualisation
- logging
- traitements pré/post-request/response

Il est très facile de définir des "**groupes de routes et sous-routeurs**" et ainsi dissocier endpoints publics et protégés (grâce aux outils `Group`).

C'était aussi rapide de gérer la **(dé-)sérialisation + validation** de requêtes et réponses (grâce aux outils `Bind` et `Render`).

