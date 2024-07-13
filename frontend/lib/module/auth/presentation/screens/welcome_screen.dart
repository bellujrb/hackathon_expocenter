import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/core/extensions/build_context_utils.dart';
import 'package:google_fonts/google_fonts.dart';

class WelcomeScreen extends StatelessWidget {
  const WelcomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    // Navega para a próxima tela após 2 segundos
    Future.delayed(const Duration(seconds: 2), () {
      // Substitua 'NextScreen' pelo nome da sua próxima tela
      Modular.to.navigate("home");
    });

    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SizedBox(
              height: 60,
              width: 218,
              child: Image.asset("assets/logo.png"),
            ),
            const SizedBox(
              height: 5,
            ),
            SizedBox(
              width: context.mediaWidth * 0.6,
              child: Text(
                "O único clube de benefício que você realmente precisa",
                style: GoogleFonts.poppins(),
                textAlign: TextAlign.center,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
