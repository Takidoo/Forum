/*ATTENTE DU CHARGEMENT INTEGRALE DE LA PAGE*/
document.addEventListener("DOMContentLoaded", function() { //on ajoute une écoute d'événement au chargement de la page
    console.log("DOM chargé"); //affiche un message console pour vérifie que le contenu est bien chargé
    
    document.getElementById("initialButtons").style.display="flex"; //on force l'affichage des boutons initiaux

    /*AFFICHAGE DYNAMIQUE AU CLIC DU BOUTON LOGIN*/
    document.getElementById("goToConnexion").addEventListener("click", function(e) {//on ajoute une écoute d'événement au clic sur le bouton Login
        e.preventDefault(); //empêche le rechargement de la page
        console.log("click sur le bouton login"); //on affiche un message console pour vérifier que le clic fonctionne
        document.getElementById("initialButtons").style.display="none"; //cache les boutons initiaux à l'activation du bouton connexion
        document.getElementById("logo").style.display="block"; //affiche le logo
        document.getElementById("connexion").style.display="flex"; //rend la div connexion visible à l'activation du bouton connexion
    });

    /*AFFICHAGE DYNAMIQUE AU CLIC DU BOUTON REGISTER*/
    document.getElementById("goToInscription").addEventListener("click", function(e) {//on ajoute une écoute d'événement au clic sur le bouton Register
        e.preventDefault(); //empêche le rechargement de la page
        console.log("click sur le bouton register"); //on affiche un message console pour vérifier que le clic fonctionne
        document.getElementById("initialButtons").style.display="none"; //rend les boutons invisibles à l'activation du bouton inscription
        document.getElementById("logo").style.display="none"; //rend le logo invisible
        document.getElementById("profilPicture").style.display="block"; //rend la div d'insertion de photo de profil visible
        document.getElementById("inscription").style.display="flex"; //rend la div inscription visible
    });

    /*PREVISUALISATION DE L'IMAGE DE PROFIL*/
    document.getElementById("picture").addEventListener("change", function(event) { //affiche l'image sélectionnée dans la div preview à partir de l'input file
        const preview = document.getElementById("preview"); //recupère la div preview avec un id unique
        const file = event.target.files[0]; //recupère le fichier sélectionné à l'indice 1 dans le tableau files
    
        if (file) { //si un fichier est sélectionné
            const reader = new FileReader(); //on crée un objet FileReader pour lire le contenu du fichier que l'on stocke dans la variable reader
            reader.onload = function(e) { //lorsque le fichier est chargé
                preview.src = e.target.result; //affiche l'image sélectionnée
                preview.style.display = "block"; //rend l'image visible
            };
            reader.readAsDataURL(file); //lit le contenu du fichier
        }
    });

    /*FONCTION QUI PERMET DE REINITIALISER LES FORMULAIRES ET LES IMAGES*/
    function resetElements(){
        //on réintialise le formulaire d'inscription
        if (document.getElementById("inscription").querySelector("form")){
            document.getElementById("inscription").querySelector("form").reset();
        }
        //on réintialise le formulaie de connexion
        if (document.getElementById("connexion").querySelector("form")){
            document.getElementById("connexion").querySelector("form").reset();
        }
        //on réinitialise l'image de profil
        document.getElementById("preview").style.display="none"; //cache l'élément preview
        document.getElementById("preview").src=""; //vide la source de l'image
            if (document.getElementById("picture")){ //réinitialise le champ input de type file (l'image posté)
                document.getElementById("picture").value="";
            }
    }


    /*AJOUT DE BOUTONS BACK DANS LES DIV CONNEXION ET INSCRIPTION*/
    //selection des conteneurs des formulaires de connexion et d'inscription
    const connexionDiv = document.getElementById("connexion"); //recupère le formulaire de connexion
    const inscriptionDiv = document.getElementById("inscription"); //recupère le formulaire d'inscription

    if (connexionDiv) { //si le formulaire de connexion existe
        console.log("formulaire de connexion trouvé, ajout du bouton Back"); //on affiche un message console
        
        //création de bouton retour au clic sur le bouton Submit du formulaire de connexion pour vérifier les identifiants et le rediriger vers la page d'accueil
        const backButtonConnexion = document.createElement("button"); //crée un bouton
        backButtonConnexion.textContent = "Back"; //ajoute le texte "Back" au bouton
        backButtonConnexion.type = "button"; //ajoute le type button au bouton
        backButtonConnexion.classList.add("backButton"); //ajoute la classe backButton au bouton

        //ajoute une écoute d'événement au clic sur le bouton retour
        backButtonConnexion.addEventListener("click", function(e) {
            e.preventDefault(); //on empêche le rechargement de la page
            console.log("click sur le bouton Back connexion"); //on affiche un message console pour vérifier que le clic du bouton retour fonctionne
            connexionDiv.style.display="none"; //on rend le formulaire invisible
            document.getElementById("logo").style.display="block"; //affiche à nouveau le logo
            document.getElementById("profilPicture").style.display="none"; //fait bien disparaitre la section avec la photo de profil
            document.getElementById("initialButtons").style.display="flex"; //les boutons initiaux visibles
            resetElements(); //on fait appel à la fonction qui réinitialise les formulaires et l'image
        });

        //on ajoute le bouton au formulaire de connexion
        connexionDiv.appendChild(backButtonConnexion) //ajoute le bouton à la div du formulaire de connexion
        }else{
            console.error("formulaire de connexion non trouvé"); //on affiche un message d'erreur sur la console si la div du formulaire de connexion n'est pas trouvé
        }

    if (inscriptionDiv) { //si le formulaire d'inscription existe
        console.log("formulaire d'inscription trouvé, ajout du bouton Back"); //on affiche un message console
        
        //on créer de bouton retour au clic sur le bouton Submit du formulaire d'inscription pour vérifier les identifiants et le rediriger vers la page d'accueil
        const backButtonInscription = document.createElement("button"); //on crée un bouton
        backButtonInscription.textContent = "Back"; //on ajoute le texte "Back" au bouton
        backButtonInscription.type = "button"; //on ajoute le type button au bouton
        backButtonInscription.classList.add("backButton"); //on ajoute la classe backButton au bouton

        //on ajoute une écoute d'événement au clic sur le bouton retour
        backButtonInscription.addEventListener("click", function(e) {
            e.preventDefault(); //on empêche le rechargement de la page
            console.log("click sur le bouton Back inscription"); //on affiche un message console pour vérifier que le clic du bouton retour fonctionne
            inscriptionDiv.style.display="none"; //on rend le formulaire invisible
            document.getElementById("profilPicture").style.display="none"; //on cache la section de photo de profil
            document.getElementById("logo").style.display="block"; //on affiche le logo à la place de la photo de profil
            document.getElementById("initialButtons").style.display="flex"; //les boutons initiaux sont visibles
            resetElements(); //on fait appel à la fonction qui réinitialise les formulaires et l'image
        });

        //on ajoute le bouton au formulaire d'inscription
        inscriptionDiv.appendChild(backButtonInscription) //on ajoute le bouton à la div du formulaire d'inscription
        }else{
            console.error("formulaire d'inscription non trouvé"); //on affiche un message d'erreur sur la console si la div du formulaire d'inscription n'est pas trouvé
        }

    /*FONCTIONNALITE DE MISE A JOUR*/
    //création de lien pour accèder à la section de mise à jour
    const forgetIdLink = document.getElementById("forgetIdlink");
    const forgetPasswordLink = document.getElementById("forgetPasswordlink");

    //mise à jour des liens
    forgetIdLink.addEventListener("click", function(e){
        e.preventDefault();
        showUpdateSection('username');
    });

    forgetPasswordLink.addEventListener("click", function(e){
        e.preventDefault();
        showUpdateSection('password');
    });

    //affiche la section de modification
    function showUpdateSection(updateType){
        //on cache les éléments suivant:
        document.getElementById("initialButtons").style.display="none";
        document.getElementById("connexion").style.display="none";
        document.getElementById("inscription").style.display="none";
        document.getElementById("profilPicture").style.display="none";
        document.getElementById("updates").style.display="none";
        //on montre les éléments suivant:
        document.getElementById("logo").style.display="block";
        document.getElementById("update").style.display="flex";
        // on cache les deux sections par défaut:
        document.getElementById("usernameUpdateData").style.display="none";
        document.getElementById("passwordUpdateData").style.display="none";
        document.getElementById("updateSubmit").style.display="block";

        //on selectionne quel type de mise à jour on veut faire (mot de passe ou id)
        if (updateType === 'username'){
            document.getElementById("usernameUpdateData").style.display="block";
            document.getElementById("passwordUpdateData").style.display="none";
        }else if(updateType === 'password'){
            document.getElementById("passwordUpdateData").style.display="block";
            document.getElementById("usernameUpdateData").style.display="none";
        }
    }
    //on ajoute des écoutes d'évenements pour les boutons de mises à jour
    document.getElementById("showUsernameUpdate").addEventListener("click", function(){
        //on montre les éléments suivant:
        document.getElementById("usernameUpdateData").style.display="block";
        document.getElementById("updateSubmit").style.display="block";
        //on cache les éléments suivant:
        document.getElementById("passwordUpdateData").style.display="none";
    });

    document.getElementById("showPasswordUpdate").addEventListener("click", function(){
        //on montre les éléments suivant:
        document.getElementById("passwordUpdateData").style.display="block";
        document.getElementById("updateSubmit").style.display="block";
        //on cache les éléments suivant:
        document.getElementById("usernameUpdateData").style.display="none";
    });

    //on créer une fonction pour gérer les messages d'erreur sur l'interface utilisateur
    function showError(inputElement, message){
        if(inputElement){ //on vérifie si l'élément existe et s'il contient une saisie
            inputElement.textContent = message;
            inputElement.style.display = "block";
        }else{
            alert(message); //si ce n'est pas le cas, on affiche une message d'erreur
        }
    }
    //on valide le formulaire de modifications
    const updateForm = document.querySelector("#update form");
    if(updateForm){
        updateForm.addEventListener("submit", function(e){
        e.preventDefault();
    
        // On vérifie l'email de l'utilisateur et on rend le champ obligatoire
        const userEmail = document.getElementById("userEmail").value.trim();
        if(!userEmail) {
            alert("Please enter your email address");
            return;
        }

        // Vérification de la validité du format de l'email
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if(!emailRegex.test(userEmail)) {
            alert("Please enter a valid email address");
            return;
        }
        
        //on vérifie que l'utilisateur a bien sélectionné une question de sécurité
        const securityQuestion = document.getElementById("securityAnswer").value.trim();
        if (!securityQuestion || securityQuestion === ""){
            alert("Please select a security question");
            return;
        }

        //si l'utilisateur n'a pas répondu à la question de sécurité:
        const securityAnswer = document.getElementById("securityAnswer").value.trim();
        if (!securityAnswer){ //si l'entrée utilisateur n'a pas été saisie
            alert("Please answer the security question"); //on déclanche une alerte en affichant un message sur l'interface utilisateur
            return;
        }
        //si l'utilisateur n'a pas entré un nouvel identifiant:
        if (document.getElementById("usernameUpdateData").style.display === "block"){ //on verifie que l'input du nouvel id est visible 
            const newUsername = document.getElementById("newUsername").value.trim(); //si c'est le cas, on récupère la valeur du champ de texte en enlevant les espaces inutiles au début et à la fin
            if (!newUsername){
                alert("Please enter a new username");
                return;
            }
        }
        
        //si l'utilisateur n'a pas entré un nouveau mot de passe:
        if (document.getElementById("passwordUpdateData").style.display === "block"){
            const newPassword = document.getElementById("newPassword").value.trim(); //on vérifie les deux saisies
            const confirmPassword = document.getElementById("confirmPassword").value.trim();

            if(!newPassword){
                alert("Please enter a new password");
                return;
            }
            if(!confirmPassword){
                alert("Please enter a new password");
                return;
            }
            if(newPassword !== confirmPassword){
                alert("Password do not match");
                return;
            }
        }
    
        //si les modifications on bien été mise à jour, on affiche un message d'alerte:
        alert("Your information has been successfully updated");

        //on réinitialise l'affichage initiale
        updateDiv.style.display="none"; //on cache la div de mise à jour
        document.getElementById("logo").style.display="block"; //on affiche le logo et les boutons initiaux
        document.getElementById("initialButtons").style.display="flex";
        document.getElementById("updates").style.display="flex";
        updateForm.reset();//on efface toutes les saisies utilisateurs
        });
    }else{
        console.error("formulaire de modification est introuvable")
    }

    /*AJOUT D'UN BOUTON BACK POUR LA MISE A JOUR UTILISATEUR*/
    const updateDiv = document.getElementById("update");
    if (updateDiv){
        const backButtonUpdate = document.createElement("button");
        backButtonUpdate.textContent = "Back";
        backButtonUpdate.type = "button";
        backButtonUpdate.classList.add("backButton");

        backButtonUpdate.addEventListener("click", function(e){
            e.preventDefault(); // empêche le refresh
            updateDiv.style.display ="none"; //cache la section de modification
            
            document.getElementById("usernameUpdateData").style.display="none"; //réinitialise les champs de saisie
            document.getElementById("passwordUpdateData").style.display="none";
            document.getElementById("updateSubmit").style.display="none";

            document.getElementById("initialButtons").style.display = "flex"; //réinitialise l'affichage initiale
            document.getElementById("logo").style.display="block";
            document.getElementById("updates").style.display="flex";

            if (updateForm){ //réinitialise le formulaire
                updateForm.reset();
            }
        });
        updateDiv.appendChild(backButtonUpdate);
    }
});