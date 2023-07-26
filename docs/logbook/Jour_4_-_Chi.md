# Jour 4 - Chi

Remplacer la config built-in du module "http" par le routeur de Chi fut quasi-immédiat, même pour les endpoints `POST`.
Pour le coup, la promesse est vraiment tenue.

J'ai commencé par faire tourner Chi.

Puis j'ai ajouté les middlewares conseillé dans la doc (RequestID, RealIP, Logger, Recoverer).

Pour les logs, il faudra remplacer les logs en texte par des logs JSON, car c'ets plus pratique pour l'exploitation dans un moteur de logs (Typesense, ELK, Grafana).

Ensuite, j'ai ajouté le middleware Chi-CORS, pour checker vite fait comment ça se passe côté sécu. Facile.
On m'a dit que Gin était mieux pourvu en termes de middlewares, etc.
Pour l'instant, je trouve que c'est fluide et je retrouve le minimum vital attendu avec Chi.

## Misc

Installer Chi :

```
go get -u github.com/go-chi/chi/v5
```

Installer Chi-CORS :

```
go get github.com/go-chi/cors
```

