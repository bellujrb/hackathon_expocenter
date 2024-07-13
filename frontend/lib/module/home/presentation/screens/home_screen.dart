import 'package:flutter/material.dart';
import 'package:frontend/core/extensions/build_context_utils.dart';
import 'package:frontend/module/home/presentation/widgets/build_appbar.dart';
import 'package:frontend/module/home/presentation/widgets/build_custom_bn.dart';
import 'package:frontend/module/home/presentation/widgets/card_fidelity.dart';
import 'package:google_fonts/google_fonts.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  int _selectedIndex = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: buildAppBar(context),
      bottomNavigationBar: CustomBottomNavigation(
        selectedIndex: _selectedIndex,
        onItemSelected: (index) {
          setState(() {
            _selectedIndex = index;
          });
        },
      ),
      body: Padding(
        padding: const EdgeInsets.all(40.0),
        child: Column(
          children: [
            SizedBox(
              height: 150,
              child: ListView.builder(
                scrollDirection: Axis.horizontal,
                itemCount: 1,
                itemBuilder: (context, index) {
                  return CardFidelity(
                    callback: (){}
                  );
                },
              ),
            ),
            const SizedBox(
              height: 20,
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  "Transações",
                  style: GoogleFonts.poppins(),
                ),
                InkWell(
                  onTap: () {},
                  child: Text(
                    "Ver todos",
                    style: GoogleFonts.poppins(),
                  ),
                ),
              ],
            ),
            const SizedBox(
              height: 10,
            ),
            SizedBox(
              height: 144,
              width: context.mediaWidth * 1.0,
              child: Image.asset("assets/payment.png"),
            )
          ],
        ),
      ),
    );
  }
}
