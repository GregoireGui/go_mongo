# go_mongo

### CI/CD
L'utilisation de la CI / CD nous permet de faire du déploiement continu de manière sécurisée grâce aux tests qui sont éxécutés (et doivent être validés) pour que le *merge* et la mise en production soit possible. Cela permet de sortir de nouvelles features et des fixs de façon plus rapide et plus sécurisée.

Nous avons choisi d'utiliser la solution CDS (Continuous Delivery Service) proposée par OVH. En effet, après étude de la [matrice de comparaison](https://github.com/ovh/cds#comparison-matrix) proposée par CDS détaillant les *features* proposées par CDS et manquantes aux autres compétiteurs, cette solution nous a semblé la plus appropriée. N'étant pas grand amateurs de la ligne de commandes, la configuration via l'interface WEB proposée par CDS a fortement penché dans la balance.

### Infrastructure

#### Base de Données
Le *free tier* proposé par MongoDB Atlas permet un stockage de 512MB sur du cloud public. Un *todo* correspond à environ 200B de données (198MB). Cela veut donc dire qu'en conservant le *free tier* de MongoDB Atlas, nous pourrions stocker environ 2 560 000 *todos*, soit 5 *todos* par jour pendant plus de 1400 ans. Compte tenu de la dimension de notre application, le *free tier* est donc largement suffisant.

Si nous voulions prendre une solution payante pour avoir plus de stockage, il existe des offres à 9$ par mois pour 2GB de stockage, et 25$ par mois pour 5GB qui seraient largement suffisantes. Toutes ces offres reposants sur des ressources partagées.

Pour des ressources allouées, le stockage commence à 10GB au prix de 57$ par mois. Dans ce cas il serait intéressant de comparer avec d'autres *providers*.

Toutefois, on observe à travers ces offres, une certaine flexibilité proposée par MongoDB Atlas.

Le cluster étant localisé en Irlande, pays membre de l'Union Européenne, nous n'avons pas à nous préoccuper des lois concernant les transferts de données hors EU, ce qui nous laisse donc comme principal inquiétude le respect du RGPD.

L'administration est simplifiée grâce à l'interface web proposée par MongoDB Atlas. En effet celle-ci nous permet facilement de visualiser nos données, nos clusters, nos projets, d'intervenir sur ceux-ci, mais aussi toutes les informations liées à nos projets. Ainsi, on peut visualiser le nombre de connections, d'opérations, le poids de nos données, mais aussi, notre facturation en direct, nos facturations passées etc. ce qui nous permet une visualisation et un contrôle total de notre cloud.

#### Serveur 
Si nous utilisons notre local comme cloud privé, les coûts d'électricité et d'entretiens seront largement inférieurs à 100 $ par mois.

Si nous devons utiliser un cloud *provider*, nous nous tournerons probablement vers la solution S.M.A.R.T Public Cloud avec ses instances s1-2 / s1-4 car il est possible de déployer des instances en Asie et en Europe à un coût très réduit (62.65$ pour 2 Web Servers (s1-2, s1-4), 1 DB server (b2-30), de l'Object Storage, et un traffic sortant gratuit de 1TB) qui est parmis les moins chers du marché. Etant donné la dimension de notre application, cette offre serait largement suffisante.

### Limitations
Concernant la base de données, les limitations concernent la puissance de la machine ainsi que la quantité de données stockées. En effet, avec le *free tier* la RAM et les vCPUs sont partagés ce qui limite notre puissance de calcul. La limitation au 512MB de stockage limite quant à elle, le nombre de *todos* pouvant être stockés. Si nous partons sur une version payante, nous pourrions investir dans des vCPUs et des RAM privés ce qui améliorerait notre puissance de calcul. La limite de stockage serait elle aussi augmentée. De plus, nous sommes limités par rapport à la géolocalisation de ce cloud car nous devons faire attention à respecter les différentes lois et réglements concernant les données européennes.

Concernant le serveur de l'application, l'utilisation de notre local induit plusieurs limitations comme la puissance de la machine en elle-même mais aussi notre infrastructure réseau. Cela peut convenir pour une petite application mais pas pour un déploiement à grande échelle.
L'utilisation d'un cloud provider nous aiderait à améliorer la puissance de calcul et aussi la qualité réseau mais avec un coût nettement supérieur, surtout si la puissance de calcul réservée est élevée. Nous pourrions cependant faire d'une pierre deux coups et n'utiliser qu'un seul provider pour notre application et le stockage de ses données et ainsi potentiellement faire baisser le coût total de l'infrastructure.