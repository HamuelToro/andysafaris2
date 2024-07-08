// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func HomeBody() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"intro__section\" id=\"home\"><div class=\"carousel\"><div class=\"list\"><div class=\"item\"><img src=\"/static/images/tourist1.jpg\" alt=\"\"><div class=\"content\"><div class=\"author\">Safaris</div><div class=\"title\">Unveils nature's breathtaking beauty</div><div class=\"topic\">Safari</div><div class=\"des\">Embark on a thrilling safari in Kenya and witness the majestic Big Five in their natural habitat. Explore the stunning Maasai Mara and immerse yourself in vibrant local cultures.</div><div class=\"buttons\"><button onclick=\"window.location.href=&#39;/safaris&#39;\">SEE MORE</button> <button href=\"./safaris.html\">BOOK NOW</button></div></div></div><div class=\"item\"><img src=\"/static/images/kenyaTaxi.jpg\" alt=\"\"><div class=\"content\"><div class=\"author\">Taxi Services</div><div class=\"title\">Do you need to a taxi??</div><div class=\"topic\">TAXI</div><div class=\"des\">Enjoy seamless airport and hotel transfers with our reliable and professional drivers. Book easily and travel stress-free to your destination.</div><div class=\"buttons\"><button onclick=\"window.location.href=&#39;/get-taxi&#39;\">SEE MORE</button> <button onclick=\"window.location.href=&#39;/get-taxi&#39;\">BOOK NOW</button></div></div></div><div class=\"item\"><img src=\"/static/images/nairobi1.jpg\" alt=\"\"><div class=\"content\"><div class=\"author\">Nairobi Tour</div><div class=\"title\">Explore The Green City In The Sun</div><div class=\"topic\">TOUR</div><div class=\"des\">Discover Nairobi's top attractions, including the Giraffe Centre and the David Sheldrick Wildlife Trust. Experience the vibrant culture and unique wildlife of Kenya's capital.</div><div class=\"buttons\"><button onclick=\"window.location.href=&#39;/nairobi-tour&#39;\">SEE MORE</button> <button onclick=\"window.location.href=&#39;/nairobi-tour/contact-form&#39;\">BOOK NOW</button></div></div></div><div class=\"item\"><img src=\"/static/images/carYard1.jpeg\" alt=\"\"><div class=\"content\"><div class=\"author\">Car Hire</div><div class=\"title\">Get safe and best deals in car hiring</div><div class=\"topic\">Car Hire</div><div class=\"des\">Hire a well-maintained car from our extensive fleet for business or leisure. Enjoy competitive rates, flexible options, and excellent customer service.</div><div class=\"buttons\"><button onclick=\"window.location.href=&#39;/car-hire&#39;\">SEE MORE</button> <button href=\"./car.html\">BOOK NOW</button></div></div></div></div><div class=\"thumbnail\"><div class=\"item\"><img src=\"/static/images/kenyaTaxi.jpg\" alt=\"\"><div class=\"content\"><div class=\"title\">Taxi Service</div></div></div><div class=\"item\"><img src=\"/static/images/nairobi1.jpg\" alt=\"\"><div class=\"content\"><div class=\"title\">Nairobi Tour</div></div></div><div class=\"item\"><img src=\"/static/images/carYard1.jpeg\" alt=\"\"><div class=\"content\"><div class=\"title\">Car Hire</div></div></div><div class=\"item\"><img src=\"/static/images/tourist1.jpg\" alt=\"\"><div class=\"content\"><div class=\"title\">Safaris</div></div></div></div><div class=\"arrows\"><button id=\"prev\"><i class=\"ri-arrow-left-s-line\"></i></button> <button id=\"next\"><i class=\"ri-arrow-right-s-line\"></i></button></div><div class=\"time\"></div></div></section><section class=\"section__get_taxi\"><div class=\"car__container\"><a href=\"#\" class=\"js-active\"><svg xmlns=\"http://www.w3.org/2000/svg\" version=\"1.1\" viewBox=\"-5.0 -10.0 110.0 135.0\"><g><path d=\"m42.781 32.957 9.1406 0.007813 2.4453 0.015625-1.1328-3.5586c-0.125-0.38672-0.21484-0.84766-0.39453-1.2109-0.11719-0.20703-0.28906-0.38281-0.50391-0.49219-0.15234-0.085938-0.32422-0.13672-0.5-0.15234-0.58594-0.039062-1.1914-0.007812-1.7852-0.011718l-3.3086-0.003907h-3.2031c-0.54297 0-1.1133-0.03125-1.6562 0.007813-0.42969 0.17969-0.73437 0.35547-0.92187 0.80469l-1.5625 4.5938c0.35547 0.058594 0.76172 0 1.1211 0z\"></path> <path d=\"m81.648 57.77c-1.1445-0.875-2.4961-1.4336-3.9258-1.6211-0.43359-0.074219-0.87891-0.097657-1.3203-0.058594l-0.015625 0.003906c-2.3945 0.23828-4.3867 1.168-5.9336 3.0586-1.3125 1.625-1.9375 3.6992-1.7422 5.7812 0.035156 0.44531 0.10156 0.89062 0.19531 1.332 0.44141 1.7773 1.4648 3.3555 2.9023 4.4883 1.582 1.2617 3.4883 1.7656 5.4922 1.7031 2.2969-0.27344 4.3398-1.2266 5.793-3.0859h0.003906c1.3633-1.7031 1.9727-3.8906 1.6836-6.0547-0.26562-2.1914-1.3906-4.1875-3.1328-5.5469zm-1.4766 9.2812c-0.75781 0.98828-1.8438 1.4844-3.0469 1.6562-1.0938 0.027344-2.1016-0.21094-2.9844-0.88672-0.94141-0.73828-1.5508-1.8164-1.6992-3-0.12891-1.1367 0.18359-2.2812 0.87109-3.1914 0.81641-1.0273 1.9023-1.5039 3.1836-1.6484 1.0039-0.078125 2.0039 0.21484 2.8047 0.82422 0.98438 0.70703 1.625 1.7969 1.7617 3.0039 0.14062 1.1562-0.17969 2.3203-0.89062 3.2461z\"></path> <path d=\"m77.512 62.562c-0.375-0.14453-0.78125-0.19141-1.1797-0.12891-0.59375 0.22656-1.0977 0.51562-1.3633 1.1289-0.20703 0.48828-0.21094 1.043-0.003906 1.5352 0.20703 0.49219 0.60547 0.87891 1.1016 1.0703 0.35547 0.14453 0.78516 0.11719 1.1562 0.074219 0.58203-0.21484 1.0352-0.51953 1.3008-1.0977v0.003907c0.21484-0.47266 0.23047-1.0039 0.050781-1.4883-0.19141-0.49219-0.57422-0.88672-1.0625-1.0977z\"></path> <path d=\"m29.09 69.457c1.4414-1.7617 2.0469-3.8594 1.8164-6.125-0.23828-2.0781-1.2773-3.9766-2.8945-5.3008-1.5742-1.2734-3.7148-2.0859-5.7617-1.8867h-0.007812c-2.1914 0.16016-4.2266 1.1875-5.6602 2.8555-1.4492 1.6992-2.043 3.7852-1.8594 6.0039s1.2656 4.1172 2.957 5.5352c1.6094 1.3398 3.6758 2.0078 5.7656 1.8555 2.2773-0.23047 4.1836-1.1562 5.6445-2.9375zm-8.7773-1.5469c-0.94141-0.66406-1.6758-1.7578-1.8594-2.8984-0.1875-1.1406 0.015625-2.2695 0.6875-3.2109 0.78906-1.0977 1.8281-1.6641 3.1523-1.8789 1.0195-0.14844 2.1758 0.18359 3.0117 0.76172 0.99219 0.69922 1.6641 1.7656 1.8633 2.9648 0.17969 1.0742-0.0625 2.1758-0.67188 3.0742-0.77344 1.1055-1.8281 1.7031-3.1484 1.9336-1.0898 0.085938-2.1328-0.10547-3.0352-0.74219z\"></path> <path d=\"m23.598 62.719c-0.33984-0.16016-0.71875-0.22266-1.0938-0.1875-0.52734 0.14062-1.0195 0.37109-1.3086 0.86719-0.26172 0.45703-0.32031 1-0.16797 1.5 0.13672 0.4375 0.42969 0.80859 0.82422 1.043 0.37109 0.20703 0.79688 0.28906 1.2188 0.23828 0.5625-0.16797 1.0078-0.40625 1.2969-0.9375v-0.003907c0.25-0.43359 0.30078-0.95312 0.14062-1.4258-0.14844-0.47266-0.47656-0.86328-0.91016-1.0938z\"></path> <path d=\"m97.16 59.879c0-1.2227 0.085938-2.5117-0.035156-3.7227-0.15625-1.8398-0.95312-3.5664-2.25-4.8789-0.95312-0.95312-2.0977-1.6055-3.3555-2.0781-1.4023-0.52734-2.9414-0.77734-4.4102-1.0469l-10.566-1.5195-2.3438-0.21484c-0.46875-0.039063-0.96094-0.054688-1.4141-0.1875l0.003906-0.003907c-0.37891-0.10937-0.73047-0.29297-1.0312-0.53906-0.49219-0.41016-0.92969-1.0039-1.3711-1.4805l-2.8047-3.0039-2.8555-3.1289c-0.37891-0.40625-0.73828-0.86328-1.1602-1.2227-0.79688-0.67187-1.6914-1.2266-2.6484-1.6445-0.94531-0.44531-1.9531-0.73828-2.9922-0.87109-1.0977-0.11719-2.2422-0.058594-3.3438-0.058594h-24.156l-0.97266 0.015625c-1.2891 0.11719-2.543 0.46094-3.6602 1.1289-2.1406 1.2812-3.8086 3.4375-5.4336 5.3008l-4.3789 5.1289-0.20312 0.23828h-7.9453c-0.58594 0-1.1484-0.007813-1.7227 0.11719-0.74609 0.16797-1.4258 0.54688-1.9648 1.082-0.80469 0.80078-1.2578 1.8906-1.2578 3.0234l0.007812 9.6055c-0.91016 0.035156-1.5352 0.13281-2.1914 0.83984-0.47656 0.54297-0.73047 1.2539-0.70312 1.9766 0.042969 0.79688 0.39844 1.5391 0.98828 2.0742 0.35547 0.32422 0.78906 0.54688 1.2578 0.65625 0.47656 0.089844 0.99219 0.066406 1.4805 0.070312h2.1406l7.8633 0.007813c-0.23828-2.5508 0.23828-5.0547 1.918-7.0625 1.6211-1.9375 3.8398-3.1641 6.3633-3.375v-0.003907c2.4375-0.20312 4.8555 0.56641 6.7305 2.1406 1.8906 1.5781 3.0625 3.8516 3.2578 6.3047 0.054688 0.66016 0.027344 1.3242-0.078125 1.9766l35.762 0.011718h-0.003906c-0.10156-0.92969-0.085938-1.8711 0.050781-2.8008 0.5-2.5859 1.8594-4.6992 4.0469-6.1797h-0.003906c4.1094-2.7383 9.6562-1.6602 12.438 2.4219 1.3672 2.0352 1.7305 4.1406 1.6094 6.5508h10.402c0.48438 0 1 0.027344 1.4727-0.058594 1.3203-0.28516 2.2695-1.4375 2.3008-2.7891 0.023437-0.76953-0.27734-1.5156-0.83203-2.0508-0.59375-0.58594-1.1992-0.72656-2.0039-0.75zm-88.312-5.832c-0.36719 0.58203-0.91406 0.86719-1.5664 1.0117-0.51953 0-1-0.0625-1.4297-0.37891-0.48047-0.32422-0.80469-0.83594-0.89453-1.4141-0.089843-0.57422 0.0625-1.1602 0.41797-1.6211 0.40234-0.51562 0.91016-0.74219 1.5469-0.82422 0.42969 0.011719 0.81641 0.058594 1.1914 0.27734 0.5 0.30469 0.85938 0.79297 1 1.3594 0.125 0.54297 0.027344 1.1172-0.26562 1.5898zm5.3984-0.14062c-0.29297 0.57812-0.82031 0.90234-1.4219 1.0977-0.47266 0.09375-0.95703 0.03125-1.3906-0.17969-0.47266-0.21094-0.84375-0.60938-1.0234-1.0977-0.20703-0.57031-0.18359-1.1953 0.066406-1.7461 0.29688-0.60156 0.84766-0.91016 1.4609-1.1133 0.43359-0.046876 0.87891-0.039063 1.2852 0.14844 0.51953 0.24609 0.92188 0.69141 1.1133 1.2344 0.19141 0.54297 0.15625 1.1406-0.09375 1.6562zm31.312-16.914 12.23 0.03125c1.2695 0.082031 2.5078 0.64844 3.5273 1.3867 1.1445 0.83203 2.082 2.0938 3.043 3.1367l4.457 4.8672-0.050781 0.085938h-12.84l-8.2109 0.042969h-1.5859c-0.16016 0-0.4375 0.035156-0.57812-0.027344zm-17.715 2.8516c1.5469-1.5742 3.3516-2.6055 5.5703-2.8047 0.6875-0.0625 1.4023-0.03125 2.0938-0.035156h7.4102v9.5078l-11.668-0.007813-7.7227-0.015625c1.3398-2.293 2.4219-4.7148 4.3164-6.6445zm14.453 14.406-4.6875 0.015625v-1.7773h4.6992zm9.1367 0.03125h-4.6055v-1.7734l4.6055 0.007813zm23.391-2.918h-55.695v-2.0859l55.695-0.007813zm18.414 6.043c-0.042969 0.16797-0.12109 0.32812-0.23828 0.46094-0.20703 0.24609-0.44141 0.29687-0.74219 0.32812h-1.7539c-0.875 0-1.6445 0.023438-2.3945-0.51953-0.55859-0.40625-1.0273-1.1133-1.125-1.8047-0.03125-0.36328-0.035156-0.72656-0.019531-1.0898 0-0.55078-0.11719-1.3984 0.23828-1.8438 0.23047-0.28906 0.51953-0.38281 0.875-0.42188h2.5469c0.54688 0 1.2617-0.070313 1.7773 0.10938l0.003906 0.003906c0.13672 0.039063 0.26562 0.10156 0.38281 0.17969 0.17188 0.125 0.42969 0.40625 0.45313 0.625 0.066406 0.625-0.007813 1.332-0.007813 1.9648 0.003906 0.63281 0.085937 1.3828 0.003906 2.0078z\"></path></g> <text x=\"25\" y=\"110\" class=\"svg-text\">TAXI</text></svg></a> <a href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" version=\"1.1\" x=\"0px\" y=\"0px\" viewBox=\"0 0 100 125\" style=\"enable-background: new 0 0 100 100\" xml:space=\"preserve\"><path d=\"M99.2,60.2c-0.3-0.2-0.1-1.1-0.1-1.1s0.7-4.5,0.6-4.9c0-0.4-1.4-2-1.4-2s1.3-3.9,0.3-4.8c-0.9-0.9-4.9-3.1-7.6-3.4  s-17.6-1.8-17.6-1.8S58.8,31.6,46.5,31c-12.3-0.6-24.7,0-24.7,0s-4.3,0.1-11.7,4.4l-6.9,5.6c0,0-0.1,0.1-0.1,0.1  c-0.1,0.1-0.2,0.2-0.3,0.3c0.1-0.1,0.1-0.1,0,0c-0.1,0.2-0.3,0.4-0.4,0.7c-0.1,0.1-0.1,0.2-0.1,0.3c0,0,0,0,0,0c0,0.1,0,0.1-0.1,0.2  c-0.1,0.2-0.1,0.3-0.1,0.5v5.3c0,0,0,0,0,0c0,0.2,0,0.5-0.1,0.7c-0.3,0.3-0.6,0.5-0.6,0.5s-1.1,0.6-1.1,1.1c0,0-0.3,5.6,0.3,7.3  c0,0,0,0,0,0c0,0,0,0,0,0c0,0.1,0.1,0.3,0.2,0.3c0,0,0,0,0,0c0.5,0.9,1.7,1.9,4.3,2.3c0,0,2.3,0.8,3.7,0.8c0.1,0,0.3,0,0.5,0  c0.1-4.7,4-8.5,8.8-8.5c4.9,0,8.8,3.9,8.8,8.8c0,0.5-0.1,1.1-0.1,1.6c0.7,0.1,1.2,0.1,1.2,0.1l44.5,0.4c-0.2-0.7-0.3-1.4-0.3-2.1  c0-4.9,3.9-8.8,8.8-8.8c4.9,0,8.8,3.9,8.8,8.8c0,0.4,0,0.9-0.1,1.3l4-0.3c0,0,5.1-0.5,5.8-1.9C99.4,60.7,99.5,60.4,99.2,60.2z   M20.4,36.1c0,0,1.5-2.1,3.2-2.3c1.8-0.2,5.3-0.4,5.3-0.4l0.4,7.1L20.4,36.1z M30.3,40.9l-0.6-7.7l11.4-0.1l0.8,8.2L30.3,40.9z   M46,41.6l-2.1-8.4c0,0,14.4-0.4,23.5,9.6L46,41.6z M97.4,49.3c-0.2,0.3-1.7,1.5-2,1.6c-0.3,0.1-1.5,0.2-2.2,0  c-0.8-0.2-1.6-0.7-1.7-0.8c-0.1-0.1-3.3-3.1-3.3-3.1s7.9,0.9,9.2,1.4C97.4,48.3,97.6,49,97.4,49.3z\"></path> <g><circle cx=\"81\" cy=\"61.5\" r=\"7.8\"></circle> <circle cx=\"18.2\" cy=\"61.5\" r=\"7.8\"></circle></g> <text x=\"30\" y=\"110\" class=\"svg-text\">SUV</text></svg></a> <a href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 20\" x=\"0px\" y=\"0px\"><path fill-rule=\"evenodd\" d=\"M11.7209302,9.22093023 C11.7209302,8.49534884 12.2976744,7.91860465 13.0232558,7.91860465 C13.3767442,7.91860465 13.6930233,8.04883721 13.9534884,8.29069767 C14.1953488,8.53255814 14.3255814,8.86744186 14.3255814,9.22093023 C14.3255814,9.94651163 13.7488372,10.5232558 13.0232558,10.5232558 C12.6697674,10.5232558 12.3534884,10.3930233 12.0930233,10.1511628 C11.8511628,9.89069767 11.7209302,9.5744186 11.7209302,9.22093023 L11.7209302,9.22093023 Z M13.2837209,6.98837209 C11.8697674,6.69069767 9.74883721,6.61627907 9.74883721,6.61627907 C9.74883721,6.61627907 8.7255814,5.87209302 7.60930233,5.5 L7.49767442,5.5744186 L7.36744186,5.74186047 C7.36744186,5.74186047 8.33488372,6.11395349 8.68837209,6.76511628 L7.72093023,6.78372093 L7.72093023,6.52325581 L7.79534884,6.3 C7.81395349,6.24418605 7.77674419,6.18837209 7.73953488,6.16976744 C7.68372093,6.15116279 7.62790698,6.18837209 7.60930233,6.24418605 L7.42325581,6.80232558 L5.41395349,6.80232558 C5.41395349,6.80232558 5.09767442,6.2627907 4.78139535,6.24418605 L4.70697674,6.24418605 C4.55813953,5.96511628 4.33488372,5.5744186 4.33488372,5.5744186 C4.33488372,5.5744186 4.29767442,5.5 4.22325581,5.5 L4.14883721,5.5 C4.14883721,5.5 4.05581395,5.5 4.09302326,5.63023256 L4.14883721,5.83488372 C3.29302326,5.5744186 2.3255814,5.29534884 2.12093023,5.74186047 C2.21395349,5.72325581 2.30697674,5.72325581 2.4372093,5.72325581 C3.03255814,5.72325581 3.77674419,5.94651163 4.24186047,6.09534884 L4.31627907,6.3372093 C3.81395349,6.05813953 1.82325581,5.66744186 1.86046512,6.18837209 C1.95348837,6.16976744 2.04651163,6.16976744 2.17674419,6.16976744 C2.86511628,6.16976744 3.68372093,6.2627907 4.39069767,6.57906977 L4.46511628,6.80232558 C4.46511628,6.80232558 3.77674419,6.39302326 2.4,6.39302326 C1.20930233,6.39302326 0.688372093,6.63488372 0.186046512,7.60232558 C0.111627907,7.75116279 -0.0558139535,8.29069767 0.0186046512,8.71860465 C0.0744186047,9.10930233 0.186046512,9.5372093 0.353488372,9.5372093 L0.520930233,9.5372093 L0.781395349,9.5372093 C0.762790698,9.44418605 0.744186047,9.33255814 0.744186047,9.22093023 C0.744186047,8.29069767 1.46976744,7.54651163 2.41860465,7.54651163 C2.86511628,7.54651163 3.2744186,7.71395349 3.60930233,8.01162791 C3.9255814,8.32790698 4.09302326,8.75581395 4.09302326,9.22093023 C4.09302326,9.33255814 4.0744186,9.4627907 4.05581395,9.5744186 L4.55813953,9.5744186 L11.0139535,9.5744186 L11.3860465,9.5744186 C11.3674419,9.4627907 11.3488372,9.35116279 11.3488372,9.22093023 C11.3488372,8.29069767 12.0744186,7.54651163 13.0232558,7.54651163 C13.4697674,7.54651163 13.8790698,7.71395349 14.1953488,8.01162791 C14.5116279,8.32790698 14.6790698,8.75581395 14.6790698,9.22093023 C14.6790698,9.35116279 14.6604651,9.4627907 14.6418605,9.59302326 L15.0325581,9.59302326 C15.0325581,9.59302326 15.8883721,9.22093023 15.9813953,8.58837209 C16.0930233,7.91860465 14.6976744,7.26744186 13.2837209,6.98837209 Z M1.11627907,9.22093023 C1.11627907,8.49534884 1.69302326,7.91860465 2.41860465,7.91860465 C2.77209302,7.91860465 3.08837209,8.04883721 3.34883721,8.29069767 C3.59069767,8.53255814 3.72093023,8.86744186 3.72093023,9.22093023 C3.72093023,9.94651163 3.14418605,10.5232558 2.41860465,10.5232558 C2.06511628,10.5232558 1.74883721,10.3930233 1.48837209,10.1511628 C1.24651163,9.89069767 1.11627907,9.5744186 1.11627907,9.22093023 L1.11627907,9.22093023 Z\"></path> <text x=\"0\" y=\"18\" class=\"svg-text vip-car\">VIP</text></svg></a> <a href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" version=\"1.1\" x=\"0px\" y=\"0px\" viewBox=\"0 0 90 62.5\" enable-background=\"new 0 0 90 50\" xml:space=\"preserve\"><g><path d=\"M71.966,32.817c-3.538,0-6.409,2.865-6.409,6.406c0,3.546,2.871,6.411,6.409,6.411s6.411-2.865,6.411-6.411   C78.377,35.683,75.504,32.817,71.966,32.817z M71.966,42.701c-1.916,0-3.472-1.553-3.472-3.475c0-1.919,1.556-3.472,3.472-3.472   s3.472,1.553,3.472,3.472C75.438,41.148,73.882,42.701,71.966,42.701z\"></path> <path d=\"M18.105,32.817c-3.538,0-6.409,2.865-6.409,6.406c0,3.546,2.871,6.411,6.409,6.411s6.411-2.865,6.411-6.411   C24.516,35.683,21.643,32.817,18.105,32.817z M18.105,42.701c-1.915,0-3.472-1.553-3.472-3.475c0-1.919,1.557-3.472,3.472-3.472   c1.917,0,3.472,1.553,3.472,3.472C21.577,41.148,20.021,42.701,18.105,42.701z\"></path> <path d=\"M88.411,28.91c-0.252-1.502-1.648-4.701-8.688-7.117l-9.847-2.742c0,0-10.542-9.141-19.703-10.187h-38.25   c0,0-8.055-0.367-9.958,14.704c0,0-1.403,11.878,1.878,14.571c0.523,0.7,1.284,1.151,2.142,1.151h4.724   c-0.009-0.15-0.023-0.301-0.023-0.453c0-4.111,3.333-7.444,7.445-7.444s7.445,3.333,7.445,7.444c0,0.152-0.014,0.303-0.023,0.453   H64.57c-0.009-0.15-0.022-0.301-0.022-0.453c0-4.111,3.333-7.444,7.444-7.444s7.444,3.333,7.444,7.444   c0,0.152-0.014,0.303-0.022,0.453h6.732v-0.062c1.138-0.253,2.038-1.289,2.264-2.614h0.056V28.91H88.411z M44.923,12.584   l-1.133,7.998l-17.97-1.237l1.02-7.199L44.923,12.584z M12.035,18.397c-2.014-0.392-1.79-2.089-1.79-2.089   c0.896-4.832,5.146-4.439,5.146-4.439l9.542,0.231l-1.008,7.116L12.035,18.397z M63.277,21.923l-17.593-1.21l1.145-8.082   l0.895,0.021c5.706-0.13,13.65,4.31,16.447,6.92S63.277,21.923,63.277,21.923z\"></path></g> <text x=\"20\" y=\"60\" class=\"svg-text mini-van\">MiniVan</text></svg></a></div><form action=\"#\" class=\"contact-form\"><div class=\"input-container\"><i class=\"ri-map-pin-line input-icon\"></i> <input type=\"text\" class=\"input-field\" name=\"address\" id=\"pickup-address\" placeholder=\"From Address\"></div><div class=\"input-container\"><i class=\"ri-map-pin-line input-icon\"></i> <input type=\"text\" class=\"input-field\" name=\"address\" id=\"dropoff-address\" placeholder=\"To...\"></div><div class=\"input-container\"><i class=\"ri-phone-fill input-icon\"></i> <input type=\"text\" class=\"input-field\" placeholder=\"Phone Number\" required></div><div class=\"input-container\"><input type=\"date\" class=\"input-field\" placeholder=\"Date and Time\" required></div></form><a class=\"get-taxi-button btn\" href=\"/get-taxi\">GET TAXI</a></section><section class=\"section__our_services\"><div class=\"service__container\"><img src=\"/static/icons/taxi.png\" alt=\"\" class=\"section-svg\"><p class=\"section-subtitle\">WHAT WE OFFER</p><p class=\"section-title\">Our Services</p><div class=\"services\"><div class=\"service__buttons\"><div class=\"service-button service-active\" data-js-data=\"taxi\">TAXI PICKUPS</div><div class=\"service-button\" data-js-data=\"nairobi\">NAIROBI TOUR</div><div class=\"service-button\" data-js-data=\"safari\">SAFARIS</div><div class=\"service-button\" data-js-data=\"hire\">CAR HIRE</div></div><div class=\"service-content show js-taxi\"><img src=\"/static/images/taxi2.jpg\" alt=\"\"><div class=\"service-des\"><p class=\"service-subtitle section-subtitle\">SERVICE IS ALWAYS THE BEST</p><p class=\"service-title section-title\">TAKE YOU ANYWHERE</p><p class=\"service-description\">Start and end your journey stress-free with our reliable airport and hotel transfer services. Whether you're traveling for business or leisure, our professional drivers ensure timely pickups and comfortable rides to your destination. Enjoy the convenience of our easy booking process and exceptional service. Book your transfer today and travel with confidence!</p><a href=\"/get-taxi\" class=\"btn book__now-button\">BOOK NOW</a></div></div><div class=\"service-content js-nairobi\"><img src=\"/static/images/nairobi2.jpg\" alt=\"\"><div class=\"service-des\"><p class=\"service-subtitle section-subtitle\">SERVICE IS ALWAYS THE BEST</p><p class=\"service-title section-title\">TOUR THE HEART OF AFRICA</p><p class=\"service-description\">Discover the vibrant heart of Kenya with our exclusive Nairobi tours. Visit iconic landmarks like the Giraffe Centre, David Sheldrick Wildlife Trust, and the Karen Blixen Museum. Experience the rich culture, stunning landscapes, and unique wildlife of Nairobi. Join us for an unforgettable adventure in Kenya's bustling capital. Book your Nairobi tour today and let the exploration begin!</p><a href=\"/nairobi-tour\" class=\"btn book__now-button\">BOOK NOW</a></div></div><div class=\"service-content js-safari\"><img src=\"/static/images/twoElephants.jpg\" alt=\"\"><div class=\"service-des\"><p class=\"service-subtitle section-subtitle\">SERVICE IS ALWAYS THE BEST</p><p class=\"service-title section-title\">NATURE'S HIDDEN WONDERs</p><p class=\"service-description\">Embark on the adventure of a lifetime with our expertly guided safaris across Kenya's breathtaking landscapes. Witness the majestic Big Five in their natural habitat, explore the iconic Maasai Mara, and immerse yourself in the vibrant local cultures. Discover the wild beauty of Kenya and create memories that will last forever. Book your safari today and let the journey begin!</p><a href=\"/safaris\" class=\"btn book__now-button\">BOOK NOW</a></div></div><div class=\"service-content js-hire\"><img src=\"/static/images/carYard2.webp\" alt=\"\"><div class=\"service-des\"><p class=\"service-subtitle section-subtitle\">SERVICE IS ALWAYS THE BEST</p><p class=\"service-title section-title\">GET CARS FROM US</p><p class=\"service-description\">Experience the freedom of the open road with our reliable car hire services. Whether you need a vehicle for business, leisure, or exploring Kenya's stunning landscapes, we offer a wide range of well-maintained cars to suit your needs. Enjoy competitive rates, flexible rental options, and exceptional customer service. Book your car hire today and drive away with ease!</p><a href=\"/car-hire\" class=\"btn book__now-button\">BOOK NOW</a></div></div></div></div></section><section class=\"call__us_section\"><div class=\"call__us_content\"><p class=\"call_us-title\">CALL US TO BOOK A TAXI</p><p class=\"call_us-number\">(+254) 700 000 000</p><p class=\"call_us-des\">Don't let travel stress ruin your trip! Call us now or book online for reliable and comfortable <span>airport and hotel transfers</span>. Enjoy peace of mind with our professional drivers and seamless service. Your journey begins with us!</p><a href=\"/get-taxi\" class=\"btn book__now-button\">BOOK NOW</a></div></section><section class=\"why__us_section\"><div class=\"why__us-content\"><img src=\"/static/icons/taxi3.png\" alt=\"\" class=\"section-svg value-section-logo\"><p class=\"section-subtitle why-us-subtitle\">WHY US</p><p class=\"section-title\">Our Core Values</p><div class=\"values-container\"><div class=\"value-container\"><i class=\"ri-thumb-up-line core-value-icon\"></i><p class=\"value-heading\">100% PLEASURE</p><p class=\"value-des\">At Andy Safaris, we prioritize your satisfaction above all else. Our dedicated team goes the extra mile to ensure every aspect of your experience is enjoyable and stress-free. From seamless booking to friendly service, we are committed to making your journey with us a delightful one. Your pleasure is our passion.</p></div><div class=\"value-container\"><i class=\"ri-planet-line core-value-icon\"></i><p class=\"value-heading\">NATIONWIDE</p><p class=\"value-des\">Our services span the entire country, offering you the flexibility to travel wherever you need. Whether you're exploring urban centers or venturing into remote destinations, we provide reliable and consistent service across Kenya. Enjoy the freedom to discover the diverse beauty of our nation with confidence, knowing we are with you every step of the way.</p></div><div class=\"value-container\"><i class=\"ri-price-tag-3-line core-value-icon\"></i><p class=\"value-heading\">FIXED PRICES</p><p class=\"value-des\">We believe in transparency and fairness, which is why we offer fixed pricing on all our services. No hidden fees, no surprises—just honest and straightforward rates. You can plan your travel budget with confidence, knowing that the price you see is the price you pay.</p></div><div class=\"value-container\"><i class=\"ri-shield-check-line core-value-icon\"></i><p class=\"value-heading\">SAFETY</p><p class=\"value-des\">Your safety is our top priority. We adhere to the highest standards of safety in all our operations, from vehicle maintenance to driver training. Our vehicles are regularly inspected and sanitized, and our drivers are trained to provide a secure and comfortable ride. Travel with peace of mind, knowing that your well-being is in good hands with us.</p></div></div></div></section><section class=\"reviews__section\"><div class=\"reviews-svg\"><img src=\"/static/icons/taxi.png\" alt=\"\" class=\"section-svg review-taxi-logo\"></div><p class=\"section-subtitle\">CUSTOMER REVIEWS</p><p class=\"section-title\">What people say about us</p><div class=\"reviews-carousel-container\"><ul class=\"reviews-carousel\" id=\"reviews-carousel\"><li class=\"review previous-review\"><div class=\"review-image\" style=\"background-image: url(/static/images/review1.jpg)\"></div><div class=\"review-text\">Booked an airport transfer and was pleasantly surprised by the professionalism and efficiency. Will definitely use their services again</div></li><li class=\"review current-review\"><div class=\"review-image\" style=\"background-image: url(/static/images/review1.jpg)\"></div><div class=\"review-text\">Great service from start to finish. The car hire was affordable and the vehicle was perfect for our needs. Highly satisfied!</div></li><li class=\"review next-review\"><div class=\"review-image\" style=\"background-image: url(/static/images/review1.jpg)\"></div><div class=\"review-text\">We used their car hire service for our trip to Nairobi. The car was in excellent condition, and the booking process was smooth and easy</div></li><li class=\"review\"><div class=\"review-image\" style=\"background-image: url(/static/images/review1.jpg)\"></div><div class=\"review-text\">Fantastic Nairobi city tour! We visited the Giraffe Centre and the David Sheldrick Wildlife Trust. The guide was excellent and made the tour very enjoyable</div></li><li class=\"review\"><div class=\"review-image\" style=\"background-image: url(/static/images/review1.jpg)\"></div><div class=\"review-text\">The team was very helpful in arranging our safari. Everything was well organized, and we had an amazing time exploring the Maasai Mara.</div></li><li class=\"review\"><div class=\"review-image\" style=\"background-image: url(/static/images/review1.jpg)\"></div><div class=\"review-text\">Their airport transfer service was top-notch. The driver was waiting for us on arrival and helped with our luggage. Very reliable!</div></li><li class=\"review\"><div class=\"review-image\" style=\"background-image: url(/static/images/review1.jpg)\"></div><div class=\"review-text\">Used their car hire service during our stay in Kenya. The process was hassle-free, and the car was in great condition. Excellent experience!\"</div></li></ul></div><div class=\"review-carousel-pips\" id=\"review-carousel-pips\"><ul class=\"list-pips\"></ul></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func HomePage(body templ.Component) templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = BaseTemplate([]string{"general", "nav", "intro", "get-taxi", "quotes", "footer", "service"}, []string{"nav", "index", "https://maps.googleapis.com/maps/api/js?key=AIzaSyD8Ij4EZwonOnh4ck32fK2EMSHkYutFTL0&libraries=places"}, body).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
