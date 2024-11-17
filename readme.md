Faire un utilitaire qui reçoit en entrée un path d'un fichier de type json , le nombre de ligne voulu par sous fichier et ensuite va découper le gros fichier en plusieurs sous fichiers plus léger. Il va vérifier que le format est bien de type json, que le premier et le dernier caractère sont une {} ou alors [] dans le cadre d'une liste de json, ou alors qu'à chaque début de ligne il y {}

Je vais séparer mon projet en plusieurs sous parties :
- Un bloc utils dans lequel il y aura mon fichier pour le traitement de split du json, et un fichier file qui aura les fonctions de vérifications et de contrôle du fichier.
- le main 

// https://medium.com/swlh/processing-16gb-file-in-seconds-go-lang-3982c235dfa2
https://www.kelche.co/blog/go/golang-bufio/