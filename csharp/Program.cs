using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Example;
using static Example.Tile.Types;

namespace ConsoleApp1
{
 
    class Program
    {
        private static void Print(Tile tile)
        {
            foreach (Layer layer in tile.Layers)
            {
                Console.WriteLine("Name : {0}", layer.Name);
                Console.WriteLine("----key----");
                for(int i=0; i<layer.Keys.Count; i++){
                    Console.WriteLine(layer.Keys[i]);                
                }
                Console.WriteLine("----Value----");
                Console.WriteLine(layer.Values.Count);
                foreach(Value value in layer.Values)
                {
                    Console.WriteLine(value.ToString());
                    Console.WriteLine(value.IntValue);
                    Console.WriteLine(value.SintValue);
                }
                foreach(Feature feature in layer.Features)
                {
                   Console.WriteLine("ID : {0}", feature.Id);
                   foreach(GeomType geo in feature.Geometry)
                   {
                     Console.WriteLine(geo.ToString());
                   }
                }
                
            }
            Console.ReadLine();
        }

       
    static void Main(string[] args)
        {
            using (Stream stream = File.OpenRead(@"C:\Users\USER\Documents\Visual Studio 2017\Projects\ConsoleApp1\ConsoleApp1\27948.pbf"))
            {
                Tile tile = Tile.Parser.ParseFrom(stream);
                Print(tile);

            }
        }
    }
}
