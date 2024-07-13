import 'package:flutter/material.dart';
import 'package:frontend/core/extensions/build_context_utils.dart';
import 'package:frontend/core/styles/colors.dart';
import 'package:google_fonts/google_fonts.dart';

class CardFidelity extends StatelessWidget {
   final VoidCallback callback;
  const CardFidelity({super.key, required this.callback});

  @override
  Widget build(BuildContext context) {
    const LinearGradient(
      colors: [
        Color(0xFF843DFF), 
        Color(0xFFFF649A), 
      ],
      begin: Alignment.topLeft,
      end: Alignment.bottomRight,
    ).createShader(const Rect.fromLTWH(0.0, 0.0, 200.0, 70.0));

    return InkWell(
      onTap: callback,
      child: Container(
        margin: const EdgeInsets.only(right: 20),
        width: context.mediaWidth * 0.8,
        decoration: BoxDecoration(
          color: AppColors.primary,
          borderRadius: BorderRadius.circular(20),
        ),
        child: Padding(
          padding: const EdgeInsets.all(24.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              Text(
                "Clube de Vantagens",
                style: GoogleFonts.poppins(
                  textStyle: const TextStyle(
                    color: Colors.white,
                    fontSize: 14
                  ),
                )
              ),
              const SizedBox(
                height: 10,
              ),
              Text(
                "Saldo",
                style: GoogleFonts.poppins(
                   textStyle: const TextStyle(
                    color: Colors.white,
                    fontSize: 12
                  ),
                ),
              ),
              Row(
                children: [
                  Text(
                    "20",
                    style: GoogleFonts.poppins(
                       textStyle: const TextStyle(
                    color: Colors.white,
                    fontWeight: FontWeight.bold,
                    fontSize: 32
                  ),
                    ),
                  ),
                  const SizedBox(
                    width: 5,
                  ),
                  const Row(
                    children: [
                      Text(
                        "\$",
                        style: TextStyle(color: Color(0xFFF8FAFC), fontSize: 16),
                      ),
                      Text(
                        "CASHBACK",
                        style: TextStyle(color: Color(0xFFF8FAFC), fontSize: 16),
                      ),
                    ],
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
