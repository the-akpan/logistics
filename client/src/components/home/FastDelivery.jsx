import Container from "../layouts/Container";

export default function FastDelivery() {
  return (
    <section className="mb-20 text-center">
      <Container>
        <div className="space-y-3 text-center">
          <h2 className="text-4xl font-bold text-slate-900">
            Ship your packages with ease
          </h2>
          <p className="text-slate-700">
            Our logistics chain lets you ship packages quickly and easily
          </p>
        </div>

        <div className="sm:w-[70%] mx-auto mt-10">
          <img
          className="w-full object-cover object-center rounded-lg"
            src={
              "https://images.unsplash.com/photo-1617347454431-f49d7ff5c3b1?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=815&q=80"
            }
            alt="Fast delivery"
          />
        </div>
      </Container>
    </section>
  );
}
