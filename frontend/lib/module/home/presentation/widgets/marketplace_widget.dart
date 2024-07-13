import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class MarketPlaceWidget extends StatelessWidget {
  final String img;
  final String name;
  final String price;
  final String token;
  final VoidCallback callback;

  const MarketPlaceWidget({
    super.key,
    required this.img,
    required this.name,
    required this.price,
    required this.token,
    required this.callback,
  });

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: callback,
      child: Container(
        height: 200,
        width: 157.5,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(20),
        ),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisAlignment: MainAxisAlignment.start,
          children: [
            ClipRRect(
              borderRadius:
                  const BorderRadius.vertical(top: Radius.circular(20)),
              child: SizedBox(
                width: double.infinity,
                child: Image.asset(
                  "assets/$img",
                  fit: BoxFit.cover,
                  alignment: Alignment.topCenter,
                ),
              ),
            ),
            const SizedBox(height: 10),
            Text(
              name,
              textAlign: TextAlign.center,
              style: GoogleFonts.poppins(
                  textStyle: const TextStyle(
                      fontSize: 16, fontWeight: FontWeight.w500),),
            ),
            const SizedBox(height: 10),
            RichText(
              text: TextSpan(
                children: [
                  TextSpan(
                    text: "\$$price ",
                    style: GoogleFonts.poppins(
                  textStyle: const TextStyle(
                    color: Colors.black,
                      fontSize: 16, fontWeight: FontWeight.w500)),
                  ),
                  TextSpan(
                    text: token,
                    style: const TextStyle(
                      color: Colors.black,
                      fontSize: 16,
                    ),
                  ),
                ],
              ),
            ),
            const SizedBox(height: 20),
          ],
        ),
      ),
    );
  }
}
