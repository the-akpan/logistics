import Container from "../layouts/Container";

export default function Testimonial() {
  return (
    <Container>
      <div class="mx-auto flex max-w-xs flex-col items-center font-bold text-slate-700 sm:max-w-xl md:max-w-3xl lg:max-w-6xl lg:flex-row">
        <div class="lg:w-[60%]">
          <h3 class="mb-8 text-xl sm:text-2xl sm:leading-[140%] lg:mb-0 lg:text-5xl">
            Get helpful,doctor-recommended health news and tips, delivered
            weekly.
          </h3>
        </div>

        <div class="lg:w-[40%]">
          <img
            className="w-full object-cover object-center rounded-lg"
            src={
              "https://images.unsplash.com/photo-1617347454431-f49d7ff5c3b1?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=815&q=80"
            }
            alt="Fast delivery"
          />
        </div>
      </div>
    </Container>
  );
}
