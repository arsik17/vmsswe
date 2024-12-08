import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/foundation.dart';

Future initFirebase() async {
  if (kIsWeb) {
    await Firebase.initializeApp(
        options: FirebaseOptions(
            apiKey: "AOSKOASKCokoKCDIOpvjiopJVODsdvdskol",
            authDomain: "vms-app-owfplb.firebaseapp.com",
            projectId: "vms-app-owfplb",
            storageBucket: "vms-app-owfplb.appspot.com",
            messagingSenderId: "257932538864",
            appId: "1:257932538864:web:da835a28a871cfafa86335"));
  } else {
    await Firebase.initializeApp();
  }
}
