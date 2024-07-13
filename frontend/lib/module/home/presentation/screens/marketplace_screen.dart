import 'package:flutter/material.dart';
import 'package:flutter_modular/flutter_modular.dart';
import 'package:frontend/core/extensions/build_context_utils.dart';
import 'package:frontend/core/styles/colors.dart';
import 'package:frontend/module/home/presentation/widgets/marketplace_widget.dart';
import 'package:google_fonts/google_fonts.dart';

class MarketPlaceScreen extends StatefulWidget {
  const MarketPlaceScreen({super.key});

  @override
  State<MarketPlaceScreen> createState() => _MarketPlaceScreenState();
}

class _MarketPlaceScreenState extends State<MarketPlaceScreen> {

  // ignore: unused_element
  void _showMarketPlaceModal() {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          backgroundColor: Colors.transparent,
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20)),
          content: Container(
            width: 335,
            height: 250,
            decoration: BoxDecoration(
              color: AppColors.primary,
              borderRadius: BorderRadius.circular(20),
            ),
            child: Column(
              children: [
                const Padding(
                  padding: EdgeInsets.symmetric(vertical: 20.0),
                  child: Icon(Icons.check_circle, size: 40, color: Colors.green),
                ),
                const Text(
                  "Beneficio coletado",
                  style: TextStyle(color: Colors.white, fontSize: 16, fontWeight: FontWeight.bold),
                ),
                const Padding(
                  padding: EdgeInsets.symmetric(vertical: 8.0),
                  child: Text(
                    "#492394294",
                    style: TextStyle(color: Colors.white, fontSize: 14),
                  ),
                ),
                const Text(
                  "Apresente este código no local.",
                  style: TextStyle(color: Colors.white, fontSize: 12),
                ),
                IconButton(
                  icon: const Icon(Icons.close, color: Colors.white),
                  onPressed: () => Modular.to.navigate("wallet"),
                ),
              ],
            ),
          ),
        );
      },
    );
  }

  final List<Map<String, String>> products = [
    {
      "img": "estacionamento.png",
      "name": "Estacionamento Expo Center Norte",
      "price": "25",
      "token": "CASHBACK"
    },
    {
      "img": "cashback.png",
      "name": "Cashback em loja especifica",
      "price": "?",
      "token": "CASHBACK"
    },
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        toolbarHeight: 40,
        title: Padding(
          padding: const EdgeInsets.all(8.0),
          child: InkWell(
            onTap: () => Modular.to.navigate("home"),
            child: SizedBox(
              height: 24,
              width: 24,
              child: Image.asset('assets/icon_arrowleft.png'),
            ),
          ),
        ),
      ),
      body: SingleChildScrollView(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(
              "Saldo",
              style: GoogleFonts.poppins(
                textStyle: const TextStyle(
                  color: Colors.black,
                  fontSize: 18,
                  fontWeight: FontWeight.w500
                )
              ),
            ),
            const SizedBox(height: 10),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  "20",
                  style: GoogleFonts.poppins(
                    textStyle: const TextStyle(
                  color: Colors.black,
                  fontSize: 24,
                  fontWeight: FontWeight.w500)
                  )
                ),
                SizedBox(width: 5),
                Text(
                  "\$CASHBACK",
                  style: TextStyle(color: Colors.black, fontSize: 16),
                ),
              ],
            ),
            const SizedBox(height: 20),
            SizedBox(
              height: context.mediaHeight * 1.0,
              child: GridView.count(
                crossAxisCount: 2,
                childAspectRatio: 0.8,
                mainAxisSpacing: 20,
                crossAxisSpacing: 20,
                padding: const EdgeInsets.all(20),
                children: products.map((product) {
                  return MarketPlaceWidget(
                    img: product["img"]!,
                    name: product["name"]!,
                    price: product["price"]!,
                    token: product["token"]!,
                    callback: _showMarketPlaceModal,
                  );
                }).toList(),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
