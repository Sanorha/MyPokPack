# __  __       _____           _      _____           _    
 |  \/  |_   _|  __ \         | |    |  __ \         | |   
 | \  / | | | | |__) |__ | | _| | __ | |__) |__ _  ___| | __
 | |\/| | | | |  ___/ _ \| |/ / |/ / |  ___/ _` |/ __| |/ /
 | |  | | |_| | |  | (_) |   <|   <  | |  | (_| | (__|   < 
 |_|  |_|\__, |_|   \___/|_|\_\_|\_\ |_|   \__,_|\___|_|\_\
          __/ |                                            
         |___/

Deviens le meilleur dresseur de la toile ! > Une application web en Go pour collectionner des Pokémon, ouvrir des boosters et voir ta collection.

🚀 Fonctionnalités
Système d'Authentification : Inscription et connexion sécurisée avec hachage de mot de passe (salt + pepper).
Ouverture de Boosters : Tente ta chance et tire 5 Pokémon aléatoires via l'API.
Ma Collection : Visualise tous les Pokémon que tu as capturés.
Interface Interactive : Une pluie de Pokéballs et de cartes rares anime ton aventure.
Sauvegarde de ta collection : Tes Pokémon et ton profil sont sauvegardés sur une base de données SQL.

🛠️ Stack Technique
Backend : Go
Frontend : HTML, JavaScript et CSS
Base de données : SQL pour la gestion des utilisateurs et des collections.

📦 Installation & Lancement
Prérequis : avoir Go d'installé sur ta machine.
Clone le dépôt.
Place-toi à la racine du projet (là où se trouve le main.go).

Lance le serveur avec la commande suivante :
go run main.go

Ensuite, ouvre ton navigateur sur : 
http://localhost:8080

📁 Structure du Projet
main.go : Point d'entrée de l'application.
donnerprojet.sqlite : base de donné SQL
index.html : Page d'accueil du site
src/ : Logique Go
pages/ : Templates et pages HTML.
static/ : Fichiers statiques CSS.

🤖 Note sur l'IA
Ce projet a été réalisé avec l'aide de l'IA pour la conception des scripts JavaScript (système de particules/objets tombants) ainsi que pour l'optimisation de certaines parties du CSS (mises en page complexes et design des cartes).
