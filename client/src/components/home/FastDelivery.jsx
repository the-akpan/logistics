import Container from "../layouts/Container";
import FeaturesSection from "./FeaturesSection";

export default function FastDelivery() {
  return (
    <section className="mb-20 text-center">
      <Container>
        <div className="space-y-3 text-center mb-8">
          <h2 className="text-4xl font-bold text-slate-900">
            Ship your packages with ease
          </h2>
          <p className="text-slate-700">
            Our logistics chain lets you ship packages quickly and easily
          </p>
        </div>
        <FeaturesSection />
      </Container>
    </section>
  );
}
