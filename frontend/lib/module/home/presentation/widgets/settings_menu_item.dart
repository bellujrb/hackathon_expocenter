import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class SettingsMenuItem extends StatelessWidget {
  final String img;
  final String title;
  final VoidCallback onTap;

  const SettingsMenuItem({
    Key? key,
    required this.title,
    required this.onTap, required this.img,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.transparent,
      child: ListTile(
        onTap: onTap,
        title: Row(
          children: [
            SizedBox(
              height: 24,
              width: 24,
              child: Image.asset("assets/$img"),
            ),
            const SizedBox(
              width: 10,
            ),
            Text(title, style: GoogleFonts.poppins()),
          ],
        ),
        trailing: const Icon(
          Icons.arrow_forward_ios,
          color: Colors.white,
          size: 16,
        ),
        contentPadding: const EdgeInsets.symmetric(horizontal: 16),
      ),
    );
  }
}
