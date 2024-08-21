// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package maintempl

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "server/api/web/layout"

func Services() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"max-w-6xl mx-auto mt-16 p-8 min-h-screen flex flex-col items-center\"><!-- Header Section --><header class=\"text-center mb-12\"><h1 class=\"text-4xl font-bold text-gray-800\">Our Services</h1><p class=\"text-lg text-gray-600 mt-4\">Explore our range of tools designed to enhance your email management experience.</p></header><!-- Services Grid --><div class=\"grid grid-cols-1 md:grid-cols-3 gap-12\"><!-- Keyword Search --><div class=\"bg-white p-8 rounded-lg shadow-xl hover:shadow-2xl transition duration-300 transform hover:scale-105\"><div class=\"flex items-center mb-4\"><svg class=\"w-12 h-12 text-indigo-500\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 4v16m8-8H4\"></path></svg><h2 class=\"text-2xl font-bold text-gray-800 ml-4\">Keyword Search</h2></div><p class=\"text-gray-600 text-lg mb-4\">Search for emails based on specific keywords with high accuracy and speed.</p><h3 class=\"text-xl font-semibold text-gray-800 mb-2\">Features:</h3><ul class=\"list-disc list-inside text-gray-600 mb-4\"><li>Advanced search algorithms</li><li>Real-time search results</li><li>Save and manage search filters</li></ul></div><!-- Email Filtering --><div class=\"bg-white p-8 rounded-lg shadow-xl hover:shadow-2xl transition duration-300 transform hover:scale-105\"><div class=\"flex items-center mb-4\"><svg class=\"w-12 h-12 text-indigo-500\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M8 7v10l9-5-9-5z\"></path></svg><h2 class=\"text-2xl font-bold text-gray-800 ml-4\">Email Filtering</h2></div><p class=\"text-gray-600 text-lg mb-4\">Filter your emails by various criteria to find exactly what you need.</p><h3 class=\"text-xl font-semibold text-gray-800 mb-2\">Capabilities:</h3><ul class=\"list-disc list-inside text-gray-600 mb-4\"><li>Customizable filter settings</li><li>Automated filtering rules</li><li>Organize emails into folders</li></ul></div><!-- Advanced Analytics --><div class=\"bg-white p-8 rounded-lg shadow-xl hover:shadow-2xl transition duration-300 transform hover:scale-105\"><div class=\"flex items-center mb-4\"><svg class=\"w-12 h-12 text-indigo-500\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 3v6h6M21 21v-6h-6M9 3h6v6H9V3zm6 12h6v6h-6v-6z\"></path></svg><h2 class=\"text-2xl font-bold text-gray-800 ml-4\">Advanced Analytics</h2></div><p class=\"text-gray-600 text-lg mb-4\">Get insights from your emails with our powerful analytics tools.</p><h3 class=\"text-xl font-semibold text-gray-800 mb-2\">Highlights:</h3><ul class=\"list-disc list-inside text-gray-600 mb-4\"><li>Data-driven insights</li><li>Customizable reports and dashboards</li><li>Integration with third-party tools</li></ul></div></div><!-- Call to Action Section --><section class=\"my-16 bg-indigo-600 text-white p-8 rounded-lg shadow-xl text-center\"><h2 class=\"text-3xl font-bold mb-4\">Ready to Optimize Your Email Experience?</h2><p class=\"text-lg mb-6\">Sign up today and start making the most of your email management with our cutting-edge tools.</p><a href=\"/signup\" class=\"inline-block bg-white text-indigo-600 font-semibold px-6 py-3 rounded hover:bg-gray-100 transition duration-300\">Sign Up Now</a></section></main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base("Services", false).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func About() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var4 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"max-w-6xl mx-auto mt-16 p-8 min-h-screen flex flex-col items-center\"><!-- Header Section --><header class=\"text-center mb-12\"><h1 class=\"text-4xl font-bold text-gray-800\">About Us</h1><p class=\"text-lg text-gray-600 mt-4\">Discover our mission, values, and the team behind Email Keyword Finder.</p></header><!-- Content Section --><div class=\"bg-white p-8 rounded-lg shadow-xl hover:shadow-2xl transition duration-300 transform hover:scale-105 w-full\"><!-- Our Mission --><section class=\"mb-12\"><h2 class=\"text-2xl font-bold text-gray-800 mb-4\">Our Mission</h2><p class=\"text-gray-600 text-lg mb-4\">At Email Keyword Finder, our mission is to make email searching and filtering as simple and efficient as possible. We strive to help individuals and businesses quickly find the information they need from their emails.</p><p class=\"text-gray-600 text-lg mb-4\">Our goal is to empower our users with the tools they need to stay organized, save time, and improve productivity. Whether you're a professional managing countless emails daily or an individual looking for a specific message, we are here to assist you.</p></section><!-- Our Team --><section class=\"mb-12\"><h2 class=\"text-2xl font-bold text-gray-800 mb-4\">Our Team</h2><p class=\"text-gray-600 text-lg mb-4\">We are a team of passionate professionals dedicated to creating tools that make your life easier. Our diverse backgrounds in technology, design, and customer service allow us to deliver top-notch products and services.</p><p class=\"text-gray-600 text-lg mb-4\">From software developers and UX designers to customer support specialists, our team collaborates closely to ensure every aspect of our product meets the highest standards.</p></section><!-- Our Values --><section class=\"mb-12\"><h2 class=\"text-2xl font-bold text-gray-800 mb-4\">Our Values</h2><p class=\"text-gray-600 text-lg mb-4\">Integrity, innovation, and customer focus are at the core of everything we do. We believe in building trust with our users by being transparent, reliable, and committed to their success.</p><p class=\"text-gray-600 text-lg mb-4\">Continuous improvement drives us to innovate and enhance our offerings, ensuring we always provide the best possible tools and services for our users.</p></section><!-- Call to Action --><section class=\"text-center mt-12\"><h2 class=\"text-3xl font-bold text-gray-800 mb-4\">Join Us on Our Journey</h2><p class=\"text-lg text-gray-600 mb-6\">Be part of a community that values efficiency, innovation, and customer success. Explore our tools and see how we can help you manage your emails better.</p><a href=\"/signup\" class=\"inline-block bg-indigo-600 text-white font-semibold px-6 py-3 rounded hover:bg-indigo-700 transition duration-300\">Get Started</a></section></div></main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base("About", false).Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Contact() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var6 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"max-w-6xl mx-auto mt-16 p-8 min-h-screen pb-24\"><!-- Header Section --><header class=\"text-center mb-12\"><h1 class=\"text-4xl font-bold text-gray-800\">Contact Us</h1><p class=\"text-lg text-gray-600 mt-4\">We'd love to hear from you! Reach out to us with any questions, feedback, or inquiries.</p></header><!-- Contact Form Section --><div class=\"bg-white p-8 rounded-lg shadow-xl hover:shadow-2xl transition duration-300 transform hover:scale-105\"><h2 class=\"text-2xl font-bold text-gray-800 mb-6\">Get in Touch</h2><form hx-post=\"/contact\" hx-target=\"body\"><div class=\"mb-6 animate__animated animate__fadeInLeft\"><label for=\"name\" class=\"block text-gray-700 font-semibold\">Name</label> <input type=\"text\" id=\"name\" name=\"name\" class=\"mt-3 p-4 w-full border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-400 focus:outline-none transition duration-300\" placeholder=\"Your Name\"></div><div class=\"mb-6 animate__animated animate__fadeInRight\"><label for=\"email\" class=\"block text-gray-700 font-semibold\">Email</label> <input type=\"email\" id=\"email\" name=\"email\" class=\"mt-3 p-4 w-full border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-400 focus:outline-none transition duration-300\" placeholder=\"Your Email\"></div><div class=\"mb-6 animate__animated animate__fadeInUp\"><label for=\"message\" class=\"block text-gray-700 font-semibold\">Message</label> <textarea id=\"message\" name=\"message\" class=\"mt-3 p-4 w-full border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-400 focus:outline-none transition duration-300\" rows=\"4\" placeholder=\"Your Message\"></textarea></div><button type=\"submit\" class=\"bg-indigo-600 text-white py-3 px-8 rounded-lg hover:bg-indigo-700 transition duration-300 transform hover:scale-105 animate__animated animate__fadeInUp\">Send Message</button></form></div><!-- Additional Contact Information Section --><section class=\"mt-16 text-center\"><h2 class=\"text-2xl font-bold text-gray-800 mb-4\">Other Ways to Reach Us</h2><p class=\"text-gray-600 text-lg mb-6\">Feel free to contact us through any of the following channels:</p><div class=\"flex flex-col md:flex-row justify-center items-center space-y-6 md:space-y-0 md:space-x-12 animate__animated animate__fadeInUp\"><div class=\"bg-indigo-100 p-6 rounded-lg shadow-md hover:shadow-lg transition duration-300 transform hover:scale-105\"><h3 class=\"text-lg font-semibold text-gray-800\">Email</h3><p class=\"text-gray-600 mt-2\">support@emailkeywordfinder.com</p></div><div class=\"bg-indigo-100 p-6 rounded-lg shadow-md hover:shadow-lg transition duration-300 transform hover:scale-105\"><h3 class=\"text-lg font-semibold text-gray-800\">Phone</h3><p class=\"text-gray-600 mt-2\">+1 (123) 456-7890</p></div><div class=\"bg-indigo-100 p-6 rounded-lg shadow-md hover:shadow-lg transition duration-300 transform hover:scale-105\"><h3 class=\"text-lg font-semibold text-gray-800\">Social Media</h3><p class=\"text-gray-600 mt-2\">Follow us on Twitter, LinkedIn, and Facebook</p></div></div></section></main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base("Contact", false).Render(templ.WithChildren(ctx, templ_7745c5c3_Var6), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}