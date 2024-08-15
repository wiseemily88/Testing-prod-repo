require 'rspec/autorun'
require_relative './calculator.rb'

describe Calculator do
  let(:calc) {Calculator.new}
  it "add two numbers" do
      expect(calc.add(3, 4)).to eq(7)
  end
  
  it "subtracts two numbers" do
      expect(calc.subtract(7,2)).to eq(5)
  end
  
  it "multiplies two numbers" do
      expect(calc.multiply(5,4)).to eq(20)
  end
  
  it "divides two numbers" do
      expect(calc.divide(30,5)).to eq(6)
  end
end
